package profile

import (
	"context"
	"fmt"

	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/apis/config"
	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/framework"
	frameworkruntime "rihib.dev/tiny-kube-scheduler/pkg/scheduler/framework/runtime"
)

func newProfile(ctx context.Context, cfg config.KubeSchedulerProfile) (framework.Framework, error) {
	return frameworkruntime.NewFramework(ctx, &cfg)
}

// Map holds frameworks indexed by scheduler name.
type Map map[string]framework.Framework

// NewMap builds the frameworks given by the configuration, indexed by name.
func NewMap(ctx context.Context, cfgs []config.KubeSchedulerProfile) (Map, error) {
	m := make(Map)
	v := cfgValidator{m: m}

	for _, cfg := range cfgs {
		p, err := newProfile(ctx, cfg)
		if err != nil {
			return nil, fmt.Errorf("creating profile for scheduler name %s: %v", cfg.SchedulerName, err)
		}
		if err := v.validate(cfg, p); err != nil {
			return nil, err
		}
		m[cfg.SchedulerName] = p
	}
	return m, nil
}

type cfgValidator struct {
	m Map
}

func (v *cfgValidator) validate(cfg config.KubeSchedulerProfile, f framework.Framework) error {
	// TODO
	return nil
}
