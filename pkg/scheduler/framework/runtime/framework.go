package runtime

import (
	"context"

	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/apis/config"
	"rihib.dev/tiny-kube-scheduler/pkg/scheduler/framework"
)

// NewFramework initializes plugins given the configuration and the registry.
func NewFramework(ctx context.Context, profile *config.KubeSchedulerProfile) (framework.Framework, error) {
	// TODO: fwkインスタンスを返す処理を行う箇所なので理解する上では重要な箇所
	return nil, nil
}
