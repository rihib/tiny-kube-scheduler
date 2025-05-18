package scheduler

import (
	"context"
	"errors"
	"fmt"

	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	configv1 "k8s.io/kube-scheduler/config/v1"
	schedulerapi "rihib.dev/tiny-kube-scheduler/pkg/scheduler/apis/config"
	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/apis/config/scheme"
	internalqueue "rihib.dev/tiny-kube-scheduler/pkg/scheduler/backend/queue"
	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/framework"
	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/profile"
)

type Scheduler struct {
	// NextPod should be a function that blocks until the next pod
	// is available. We don't use a channel for this, because scheduling
	// a pod may take some amount of time and we don't want pods to get
	// stale while they sit in a channel.
	NextPod func(logger klog.Logger) (*framework.QueuedPodInfo, error)

	// Close this to shut down the scheduler.
	StopEverything <-chan struct{}

	// SchedulingQueue holds pods to be scheduled
	SchedulingQueue internalqueue.SchedulingQueue

	// logger *must* be initialized when creating a Scheduler,
	// otherwise logging functions will access a nil sink and
	// panic.
	logger klog.Logger
}

type schedulerOptions struct {
	profiles            []schedulerapi.KubeSchedulerProfile
	applyDefaultProfile bool
}

var defaultSchedulerOptions = schedulerOptions{
	applyDefaultProfile: true,
}

// New returns a Scheduler
func New(ctx context.Context,
	client clientset.Interface,
	informerFactory informers.SharedInformerFactory) (*Scheduler, error) {

	logger := klog.FromContext(ctx)
	StopEverything := ctx.Done()

	options := defaultSchedulerOptions

	if options.applyDefaultProfile {
		var versionedCfg configv1.KubeSchedulerConfiguration
		scheme.Scheme.Default(&versionedCfg)
		cfg := schedulerapi.KubeSchedulerConfiguration{}
		if err := scheme.Scheme.Convert(&versionedCfg, &cfg, nil); err != nil { // FIXME: KubeSchedulerConfigurationをちゃんと実装してないのでconvertでエラー出る可能性あり
			return nil, err
		}
		options.profiles = cfg.Profiles
	}

	profiles, err := profile.NewMap(ctx, options.profiles)
	if err != nil {
		return nil, fmt.Errorf("initializing profiles: %v", err)
	}

	if len(profiles) == 0 {
		return nil, errors.New("at least one profile is required")
	}

	podQueue := internalqueue.NewSchedulingQueue(
		profiles[options.profiles[0].SchedulerName].QueueSortFunc(),
		informerFactory,
	)

	sched := &Scheduler{
		StopEverything:  StopEverything,
		SchedulingQueue: podQueue,
		logger:          logger,
	}
	sched.NextPod = podQueue.Pop

	if err = addAllEventHandlers(sched, informerFactory); err != nil {
		return nil, fmt.Errorf("adding event handlers: %w", err)
	}

	return sched, nil
}
