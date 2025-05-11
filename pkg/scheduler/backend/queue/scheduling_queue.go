package queue

import (
	"sync"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"rihib.dev/tiny-kube-schedulers/pkg/scheduler/backend/heap"
	"rihib.dev/tiny-kube-schedulers/pkg/scheduler/framework"
)

type SchedulingQueue interface {
	Add(pod *v1.Pod)
}

func NewSchedulingQueue(lessFn framework.LessFunc, informerFactory informers.SharedInformerFactory) SchedulingQueue {
	return NewPriorityQueue(lessFn, informerFactory)
}

type PriorityQueue struct {
	// lock takes precedence and should be taken first,
	// before any other locks in the queue (activeQueue.lock).
	// Correct locking order is: lock > activeQueue.lock.
	lock sync.RWMutex

	activeQ activeQueuer
}

func NewPriorityQueue(lessFn framework.LessFunc, informerFactory informers.SharedInformerFactory) *PriorityQueue {
	pq := &PriorityQueue{}
	pq.activeQ = newActiveQueue(heap.New(podInfoKeyFunc, heap.LessFunc[*framework.QueuedPodInfo](lessFn)))
	return pq
}

// moveToActiveQ tries to add the pod to the active queue.
func (p *PriorityQueue) moveToActiveQ(pInfo *framework.QueuedPodInfo) bool {
	added := false
	p.activeQ.underLock(func(unlockedActiveQ unlockedActiveQueuer) {
		unlockedActiveQ.add(pInfo)
		added = true
	})
	return added
}

func (p *PriorityQueue) Add(pod *v1.Pod) {
	p.lock.Lock()
	defer p.lock.Unlock()

	pInfo := p.newQueuedPodInfo(pod)
	if added := p.moveToActiveQ(pInfo); added {
		p.activeQ.broadcast()
	}
}

func (p *PriorityQueue) newQueuedPodInfo(pod *v1.Pod) *framework.QueuedPodInfo {
	podInfo, _ := framework.NewPodInfo(pod)
	return &framework.QueuedPodInfo{
		PodInfo: podInfo,
	}
}

func podInfoKeyFunc(pInfo *framework.QueuedPodInfo) string {
	return cache.NewObjectName(pInfo.Pod.Namespace, pInfo.Pod.Name).String()
}
