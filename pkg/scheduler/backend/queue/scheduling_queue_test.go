package queue

import (
	"rihib.dev/tiny-kube-schedulers/pkg/scheduler/framework"
	"rihib.dev/tiny-kube-schedulers/pkg/scheduler/framework/plugins/queuesort"
)

func newDefaultQueueSort() framework.LessFunc {
	sort := &queuesort.PrioritySort{}
	return sort.Less
}
