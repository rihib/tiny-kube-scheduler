package scheduler

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/klog/v2/ktesting"
	internalqueue "rihib.dev/tiny-kube-scheduler/pkg/scheduler/backend/queue"
	st "rihib.dev/tiny-kube-scheduler/pkg/scheduler/testing"
)

// original
func TestAddAllEventHandlers(t *testing.T) {
	tests := []struct {
		name                  string
		expectStaticInformers map[reflect.Type]bool
		expectPodName         string
	}{
		{
			name: "default handlers in framework",
			expectStaticInformers: map[reflect.Type]bool{
				reflect.TypeOf(&v1.Pod{}): true,
			},
			expectPodName: "p",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, ctx := ktesting.NewTestContext(t)
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			fakeClient := fake.NewClientset()
			informerFactory := informers.NewSharedInformerFactory(fakeClient, 0)
			schedulingQueue := internalqueue.NewTestQueueWithInformerFactory(ctx, nil, informerFactory)
			testSched := Scheduler{
				StopEverything:  ctx.Done(),
				SchedulingQueue: schedulingQueue,
				logger:          logger,
			}

			if err := addAllEventHandlers(&testSched, informerFactory); err != nil {
				t.Fatalf("Add event handlers failed, error = %v", err)
			}

			informerFactory.Start(testSched.StopEverything)
			staticInformers := informerFactory.WaitForCacheSync(testSched.StopEverything)

			if diff := cmp.Diff(tt.expectStaticInformers, staticInformers); diff != "" {
				t.Errorf("Unexpected diff (-want, +got):\n%s", diff)
			}

			pod := st.MakePod().Name(tt.expectPodName).Obj()
			if _, err := fakeClient.CoreV1().Pods("").Create(ctx, pod, metav1.CreateOptions{}); err != nil {
				t.Fatalf("Create pod failed, error = %v", err)
			}

			gotPod, err := testSched.SchedulingQueue.Pop(logger)
			if err != nil {
				t.Fatalf("pop failed, %s", err)
			}
			if diff := cmp.Diff(tt.expectPodName, gotPod.Pod.Name); diff != "" {
				t.Errorf("Unexpected pod name (-want, +got):\n%s", diff)
			}
		})
	}
}
