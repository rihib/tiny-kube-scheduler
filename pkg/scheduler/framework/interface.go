package framework

type Plugin interface {
	Name() string
}

type LessFunc func(podInfo1, podInfo2 *QueuedPodInfo) bool

type QueueSortPlugin interface {
	Plugin
	Less(*QueuedPodInfo, *QueuedPodInfo) bool
}

// Handle provides data and some tools that plugins can use. It is
// passed to the plugin factories at the time of plugin initialization. Plugins
// must store and use this handle to call framework functions.
type Handle interface{}
