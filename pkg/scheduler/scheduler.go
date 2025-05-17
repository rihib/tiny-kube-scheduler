package scheduler

import (
	"k8s.io/klog/v2"
	internalqueue "rihib.dev/tiny-kube-scheduler/pkg/scheduler/backend/queue"
)

type Scheduler struct {
	// Close this to shut down the scheduler.
	StopEverything <-chan struct{}

	// SchedulingQueue holds pods to be scheduled
	SchedulingQueue internalqueue.SchedulingQueue

	// logger *must* be initialized when creating a Scheduler,
	// otherwise logging functions will access a nil sink and
	// panic.
	logger klog.Logger
}
