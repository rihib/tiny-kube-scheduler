package queue

import (
	"testing"

	"k8s.io/klog/v2/ktesting"
	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/backend/heap"
	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/framework"
	st "rihib.dev/tiny-kube-scheduler/pkg/scheduler/testing"
)

func TestClose(t *testing.T) {
	logger, _ := ktesting.NewTestContext(t)
	aq := newActiveQueue(heap.New(podInfoKeyFunc, heap.LessFunc[*framework.QueuedPodInfo](newDefaultQueueSort())))

	aq.underLock(func(unlockedActiveQ unlockedActiveQueuer) {
		unlockedActiveQ.add(&framework.QueuedPodInfo{PodInfo: &framework.PodInfo{Pod: st.MakePod().Name("p1").Obj()}})
		unlockedActiveQ.add(&framework.QueuedPodInfo{PodInfo: &framework.PodInfo{Pod: st.MakePod().Name("p2").Obj()}})
	})

	_, err := aq.pop(logger)
	if err != nil {
		t.Fatalf("unexpected error while pop(): %v", err)
	}
	_, err = aq.pop(logger)
	if err != nil {
		t.Fatalf("unexpected error while pop(): %v", err)
	}

	aq.close()
}
