package scheduler

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

func addAllEventHandlers(sched *Scheduler, informerFactory informers.SharedInformerFactory) error {
	if _, err := informerFactory.Core().V1().Pods().Informer().AddEventHandler(
		cache.FilteringResourceEventHandler{
			FilterFunc: isUnassignedPod,
			Handler: cache.ResourceEventHandlerFuncs{
				AddFunc: sched.addPodToSchedulingQueue,
			},
		},
	); err != nil {
		return err
	}
	return nil
}

// original
func isUnassignedPod(obj interface{}) bool {
	switch t := obj.(type) {
	case *v1.Pod:
		return !assignedPod(t)
	default:
		return false
	}
}

func assignedPod(pod *v1.Pod) bool {
	return len(pod.Spec.NodeName) != 0
}

func (sched *Scheduler) addPodToSchedulingQueue(obj interface{}) {
	pod := obj.(*v1.Pod)
	sched.SchedulingQueue.Add(pod)
}
