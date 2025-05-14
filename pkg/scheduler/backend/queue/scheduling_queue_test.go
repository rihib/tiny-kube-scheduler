package queue

import (
	"context"
	"testing"

	v1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/ktesting"
	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/framework"
	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/framework/plugins/queuesort"
	st "rihib.dev/tiny-kube-scheduler/pkg/scheduler/testing"
)

// original
var (
	podInfo = mustNewPodInfo(
		st.MakePod().Name("p").Namespace("ns").UID("pns").Obj(),
	)
)

// original
func TestPriorityQueue_Add(t *testing.T) {
	logger, ctx := ktesting.NewTestContext(t)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	q := NewTestQueue(ctx, newDefaultQueueSort())
	q.Add(logger, podInfo.Pod)
	if p, err := q.Pop(logger); err != nil || p.Pod != podInfo.Pod {
		t.Errorf("Expected: %v after Pop, but got: %v", podInfo.Pod.Name, p.Pod.Name)
	}
}

func newDefaultQueueSort() framework.LessFunc {
	sort := &queuesort.PrioritySort{}
	return sort.Less
}

func (p *PriorityQueue) Pop(logger klog.Logger) (*framework.QueuedPodInfo, error) {
	return p.activeQ.pop(logger)
}

func mustNewPodInfo(pod *v1.Pod) *framework.PodInfo {
	podInfo, err := framework.NewPodInfo(pod)
	if err != nil {
		panic(err)
	}
	return podInfo
}
