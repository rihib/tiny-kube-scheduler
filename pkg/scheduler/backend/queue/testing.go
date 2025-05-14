package queue

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"
	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/framework"
)

// NewTestQueue creates a priority queue with an empty informer factory.
func NewTestQueue(ctx context.Context, lessFn framework.LessFunc) *PriorityQueue {
	return NewTestQueueWithObjects(ctx, lessFn, nil)
}

func NewTestQueueWithObjects(
	ctx context.Context,
	lessFn framework.LessFunc,
	objs []runtime.Object,
) *PriorityQueue {
	informerFactory := informers.NewSharedInformerFactory(fake.NewClientset(objs...), 0)

	return NewTestQueueWithInformerFactory(ctx, lessFn, informerFactory)
}

func NewTestQueueWithInformerFactory(
	ctx context.Context,
	lessFn framework.LessFunc,
	informerFactory informers.SharedInformerFactory,
) *PriorityQueue {
	pq := NewPriorityQueue(lessFn, informerFactory)
	informerFactory.Start(ctx.Done())
	informerFactory.WaitForCacheSync(ctx.Done())
	return pq
}
