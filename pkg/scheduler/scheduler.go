package scheduler

import (
	internalqueue "rihib.dev/tiny-kube-schedulers/pkg/scheduler/backend/queue"
)

type Scheduler struct {
	SchedulingQueue internalqueue.SchedulingQueue
}
