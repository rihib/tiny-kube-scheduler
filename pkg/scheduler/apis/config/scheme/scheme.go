package scheme

import "k8s.io/apimachinery/pkg/runtime"

var (
	// Scheme is the runtime.Scheme to which all kubescheduler api types are registered.
	Scheme = runtime.NewScheme()
)
