package queue

import (
	"sync"

	v1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/backend/heap"
	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/framework"
)

// activeQueuer is a wrapper for activeQ related operations.
type activeQueuer interface {
	underLock(func(unlockedActiveQ unlockedActiveQueuer))

	pop(logger klog.Logger) (*framework.QueuedPodInfo, error)
	list() []*v1.Pod

	close()
	broadcast()
}

// unlockedActiveQueuer defines activeQ methods that are not protected by the lock itself.
// underLock() method should be used to protect these methods.
type unlockedActiveQueuer interface {
	// add adds a new pod to the activeQ.
	// This method should be called in activeQueue.underLock().
	add(pInfo *framework.QueuedPodInfo)
}

// unlockedActiveQueue defines activeQ methods that are not protected by the lock itself.
// activeQueue.underLock() method should be used to protect these methods.
type unlockedActiveQueue struct {
	queue *heap.Heap[*framework.QueuedPodInfo]
}

func newUnlockedActiveQueue(queue *heap.Heap[*framework.QueuedPodInfo]) *unlockedActiveQueue {
	return &unlockedActiveQueue{
		queue: queue,
	}
}

func (uaq *unlockedActiveQueue) add(pInfo *framework.QueuedPodInfo) {
	uaq.queue.AddOrUpdate(pInfo)
}

type activeQueue struct {
	lock          sync.RWMutex
	queue         *heap.Heap[*framework.QueuedPodInfo]
	unlockedQueue *unlockedActiveQueue
	cond          sync.Cond
	closed        bool
}

func newActiveQueue(queue *heap.Heap[*framework.QueuedPodInfo]) *activeQueue {
	aq := &activeQueue{
		queue:         queue,
		unlockedQueue: newUnlockedActiveQueue(queue),
	}
	aq.cond.L = &aq.lock

	return aq
}

func (aq *activeQueue) underLock(fn func(unlockedActiveQ unlockedActiveQueuer)) {
	aq.lock.Lock()
	defer aq.lock.Unlock()
	fn(aq.unlockedQueue)
}

func (aq *activeQueue) pop(logger klog.Logger) (*framework.QueuedPodInfo, error) {
	aq.lock.Lock()
	defer aq.lock.Unlock()

	return aq.unlockedPop(logger)
}

func (aq *activeQueue) unlockedPop(logger klog.Logger) (*framework.QueuedPodInfo, error) {
	var pInfo *framework.QueuedPodInfo
	for aq.queue.Len() == 0 {
		if aq.closed {
			return nil, nil
		}
		aq.cond.Wait()
	}
	pInfo, err := aq.queue.Pop()
	if err != nil {
		return nil, err
	}

	return pInfo, nil
}

func (aq *activeQueue) list() []*v1.Pod {
	aq.lock.RLock()
	defer aq.lock.RUnlock()
	var result []*v1.Pod
	for _, pInfo := range aq.queue.List() {
		result = append(result, pInfo.Pod)
	}
	return result
}

func (aq *activeQueue) close() {
	aq.lock.Lock()
	defer aq.lock.Unlock()
	aq.closed = true
}

// broadcast notifies the pop() operation that new pod(s) was added to the activeQueue.
func (aq *activeQueue) broadcast() {
	aq.cond.Broadcast()
}
