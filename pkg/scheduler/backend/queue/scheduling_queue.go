package queue

import (
	"sync"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	"k8s.io/utils/clock"
	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/backend/heap"
	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/framework"
)

type SchedulingQueue interface {
	Add(logger klog.Logger, pod *v1.Pod)
}

func NewSchedulingQueue(lessFn framework.LessFunc, informerFactory informers.SharedInformerFactory) SchedulingQueue {
	return NewPriorityQueue(lessFn, informerFactory)
}

type PriorityQueue struct {
	clock clock.WithTicker

	// lock takes precedence and should be taken first,
	// before any other locks in the queue (activeQueue.lock).
	// Correct locking order is: lock > activeQueue.lock.
	lock sync.RWMutex

	activeQ activeQueuer
}

var defaultPriorityQueueOptions = priorityQueueOptions{
	clock: clock.RealClock{},
}

type priorityQueueOptions struct {
	clock clock.WithTicker
}

func NewPriorityQueue(lessFn framework.LessFunc, informerFactory informers.SharedInformerFactory) *PriorityQueue {
	options := defaultPriorityQueueOptions

	pq := &PriorityQueue{
		clock: options.clock,
	}
	pq.activeQ = newActiveQueue(heap.New(podInfoKeyFunc, heap.LessFunc[*framework.QueuedPodInfo](lessFn)))
	return pq
}

// moveToActiveQ tries to add the pod to the active queue.
func (p *PriorityQueue) moveToActiveQ(_ klog.Logger, pInfo *framework.QueuedPodInfo) bool {
	added := false
	p.activeQ.underLock(func(unlockedActiveQ unlockedActiveQueuer) {
		if pInfo.InitialAttemptTimestamp == nil {
			now := p.clock.Now()
			pInfo.InitialAttemptTimestamp = &now
		}

		unlockedActiveQ.add(pInfo)
		added = true
	})
	return added
}

func (p *PriorityQueue) Add(logger klog.Logger, pod *v1.Pod) {
	p.lock.Lock()
	defer p.lock.Unlock()

	pInfo := p.newQueuedPodInfo(pod)
	if added := p.moveToActiveQ(logger, pInfo); added {
		p.activeQ.broadcast()
	}
}

func (p *PriorityQueue) newQueuedPodInfo(pod *v1.Pod) *framework.QueuedPodInfo {
	now := p.clock.Now()
	// ignore this err since apiserver doesn't properly validate affinity terms
	// and we can't fix the validation for backwards compatibility.
	podInfo, _ := framework.NewPodInfo(pod)
	return &framework.QueuedPodInfo{
		PodInfo:                 podInfo,
		Timestamp:               now,
		InitialAttemptTimestamp: nil,
	}
}

func podInfoKeyFunc(pInfo *framework.QueuedPodInfo) string {
	return cache.NewObjectName(pInfo.Pod.Namespace, pInfo.Pod.Name).String()
}
