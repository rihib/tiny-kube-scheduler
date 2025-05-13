package scheduler

import (
	internalqueue "rihib.dev/tiny-kube-scheduler/pkg/scheduler/backend/queue"
)

type Scheduler struct {
	SchedulingQueue internalqueue.SchedulingQueue
}
