package queuesort

import (
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"
	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/framework"
	st "rihib.dev/tiny-kube-scheduler/pkg/scheduler/testing"
)

func TestLess(t *testing.T) {
	prioritySort := &PrioritySort{}
	var lowPriority, highPriority = int32(10), int32(100)
	t1 := time.Now()
	t2 := t1.Add(time.Second)
	for _, tt := range []struct {
		name     string
		p1       *framework.QueuedPodInfo
		p2       *framework.QueuedPodInfo
		expected bool
	}{
		{
			name: "p1.priority less than p2.priority",
			p1: &framework.QueuedPodInfo{
				PodInfo: mustNewPodInfo(t, st.MakePod().Priority(lowPriority).Obj()),
			},
			p2: &framework.QueuedPodInfo{
				PodInfo: mustNewPodInfo(t, st.MakePod().Priority(highPriority).Obj()),
			},
			expected: false, // p2 should be ahead of p1 in the queue
		},
		{
			name: "p1.priority greater than p2.priority",
			p1: &framework.QueuedPodInfo{
				PodInfo: mustNewPodInfo(t, st.MakePod().Priority(highPriority).Obj()),
			},
			p2: &framework.QueuedPodInfo{
				PodInfo: mustNewPodInfo(t, st.MakePod().Priority(lowPriority).Obj()),
			},
			expected: true, // p1 should be ahead of p2 in the queue
		},
		{
			name: "equal priority. p1 is added to schedulingQ earlier than p2",
			p1: &framework.QueuedPodInfo{
				PodInfo:   mustNewPodInfo(t, st.MakePod().Priority(highPriority).Obj()),
				Timestamp: t1,
			},
			p2: &framework.QueuedPodInfo{
				PodInfo:   mustNewPodInfo(t, st.MakePod().Priority(highPriority).Obj()),
				Timestamp: t2,
			},
			expected: true, // p1 should be ahead of p2 in the queue
		},
		{
			name: "equal priority. p2 is added to schedulingQ earlier than p1",
			p1: &framework.QueuedPodInfo{
				PodInfo:   mustNewPodInfo(t, st.MakePod().Priority(highPriority).Obj()),
				Timestamp: t2,
			},
			p2: &framework.QueuedPodInfo{
				PodInfo:   mustNewPodInfo(t, st.MakePod().Priority(highPriority).Obj()),
				Timestamp: t1,
			},
			expected: false, // p2 should be ahead of p1 in the queue
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := prioritySort.Less(tt.p1, tt.p2); got != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}

func mustNewPodInfo(t *testing.T, pod *v1.Pod) *framework.PodInfo {
	podInfo, err := framework.NewPodInfo(pod)
	if err != nil {
		t.Fatal(err)
	}
	return podInfo
}
