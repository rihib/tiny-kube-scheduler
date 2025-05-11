package framework

import v1 "k8s.io/api/core/v1"

type QueuedPodInfo struct {
	*PodInfo
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
