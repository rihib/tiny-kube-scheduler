package framework

import (
	"time"

	v1 "k8s.io/api/core/v1"
)

type QueuedPodInfo struct {
	*PodInfo
	// The time pod added to the scheduling queue.
	Timestamp time.Time
	// The time when the pod is added to the queue for the first time. The pod may be added
	// back to the queue multiple times before it's successfully scheduled.
	// It shouldn't be updated once initialized. It's used to record the e2e scheduling
	// latency for a pod.
	InitialAttemptTimestamp *time.Time
}

type PodInfo struct {
	Pod *v1.Pod
}

func (pi *PodInfo) Update(pod *v1.Pod) error {
	pi.Pod = pod
	return nil
}

func NewPodInfo(pod *v1.Pod) (*PodInfo, error) {
	pInfo := &PodInfo{}
	err := pInfo.Update(pod)
	return pInfo, err
}
