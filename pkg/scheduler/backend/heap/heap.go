package heap

import (
	"container/heap"
	"fmt"
)

// KeyFunc is a function type to get the key from an object.
type KeyFunc[T any] func(obj T) string

type heapItem[T any] struct {
	obj   T   // The object which is stored in the heap.
	index int // The index of the object's key in the Heap.queue.
}

type itemKeyValue[T any] struct {
	key string
	obj T
}

// data is an internal struct that implements the standard heap interface
// and keeps the data stored in the heap.
type data[T any] struct {
	// items is a map from key of the objects to the objects and their index.
	// We depend on the property that items in the map are in the queue and vice versa.
	items map[string]*heapItem[T]
	// queue implements a heap data structure and keeps the order of elements
	// according to the heap invariant. The queue keeps the keys of objects stored
	// in "items".
	queue []string

	// keyFunc is used to make the key used for queued item insertion and retrieval, and
	// should be deterministic.
	keyFunc KeyFunc[T]
	// lessFunc is used to compare two objects in the heap.
	lessFunc LessFunc[T]
}

// Less compares two objects and returns true if the first one should go
// in front of the second one in the heap.
func (h *data[T]) Less(i, j int) bool {
	if i > len(h.queue) || j > len(h.queue) {
		return false
	}
	itemi, ok := h.items[h.queue[i]]
	if !ok {
		return false
	}
	itemj, ok := h.items[h.queue[j]]
	if !ok {
		return false
	}
	return h.lessFunc(itemi.obj, itemj.obj)
}

func (h *data[T]) Len() int { return len(h.queue) }

func (h *data[T]) Swap(i, j int) {
	if i < 0 || j < 0 {
		return
	}
	h.queue[i], h.queue[j] = h.queue[j], h.queue[i]
	item := h.items[h.queue[i]]
	item.index = i
	item = h.items[h.queue[j]]
	item.index = j
}

func (h *data[T]) Push(kv interface{}) {
	keyValue := kv.(*itemKeyValue[T])
	n := len(h.queue)
	h.items[keyValue.key] = &heapItem[T]{keyValue.obj, n}
	h.queue = append(h.queue, keyValue.key)
}

func (h *data[T]) Pop() interface{} {
	if len(h.queue) == 0 {
		return nil
	}
	key := h.queue[len(h.queue)-1]
	h.queue = h.queue[0 : len(h.queue)-1]
	item, ok := h.items[key]
	if !ok {
		// This is an error
		return nil
	}
	delete(h.items, key)
	return item.obj
}

func (h *data[T]) Peek() (T, bool) {
	if len(h.queue) > 0 {
		return h.items[h.queue[0]].obj, true
	}
	var zero T
	return zero, false
}

// Heap is a producer/consumer queue that implements a heap data structure.
// It can be used to implement priority queues and similar data structures.
type Heap[T any] struct {
	// data stores objects and has a queue that keeps their ordering according
	// to the heap invariant.
	data *data[T]
}

func (h *Heap[T]) Delete(obj T) error {
	key := h.data.keyFunc(obj)
	if item, ok := h.data.items[key]; ok {
		heap.Remove(h.data, item.index)
		return nil
	}
	return fmt.Errorf("object not found")
}

func (h *Heap[T]) Peek() (T, bool) {
	return h.data.Peek()
}

func (h *Heap[T]) Pop() (T, error) {
	obj := heap.Pop(h.data)
	if obj != nil {
		return obj.(T), nil
	}
	var zero T
	return zero, fmt.Errorf("heap is empty")
}

// AddOrUpdate inserts an item, and puts it in the queue. The item is updated if it
// already exists.
func (h *Heap[T]) AddOrUpdate(obj T) {
	key := h.data.keyFunc(obj)
	if _, exists := h.data.items[key]; exists { // TODO: replace exists with ok
		h.data.items[key].obj = obj
		heap.Fix(h.data, h.data.items[key].index)
	} else {
		heap.Push(h.data, &itemKeyValue[T]{key, obj})
	}
}

func (h *Heap[T]) Len() int {
	return len(h.data.queue)
}

func New[T any](keyFn KeyFunc[T], lessFn LessFunc[T]) *Heap[T] {
	return &Heap[T]{
		data: &data[T]{
			items:    map[string]*heapItem[T]{},
			queue:    []string{},
			keyFunc:  keyFn,
			lessFunc: lessFn,
		},
	}
}

// LessFunc is a function that receives two items and returns true if the first
// item should be placed before the second one when the list is sorted.
type LessFunc[T any] func(item1, item2 T) bool
