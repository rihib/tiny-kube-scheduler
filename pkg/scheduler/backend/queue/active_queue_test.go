package queue

import (
	"testing"

	"rihib.dev/tiny-kube-schedulers/pkg/scheduler/backend/heap"
	"rihib.dev/tiny-kube-schedulers/pkg/scheduler/framework"
	st "rihib.dev/tiny-kube-schedulers/pkg/scheduler/testing"
)

func TestClose(t *testing.T) {
	aq := newActiveQueue(heap.New(podInfoKeyFunc, heap.LessFunc[*framework.QueuedPodInfo](newDefaultQueueSort())))

	aq.underLock(func(unlockedActiveQ unlockedActiveQueuer) {
		unlockedActiveQ.add(&framework.QueuedPodInfo{PodInfo: &framework.PodInfo{Pod: st.MakePod().Namespace("foo").Name("p1").UID("p1").Obj()}})
		unlockedActiveQ.add(&framework.QueuedPodInfo{PodInfo: &framework.PodInfo{Pod: st.MakePod().Namespace("bar").Name("p2").UID("p2").Obj()}})
	})

	_, err := aq.pop()
	if err != nil {
		t.Fatalf("unexpected error while pop(): %v", err)
	}
	_, err = aq.pop()
	if err != nil {
		t.Fatalf("unexpected error while pop(): %v", err)
	}

	aq.close()
}
