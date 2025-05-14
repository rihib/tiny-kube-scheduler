module rihib.dev/tiny-kube-scheduler

go 1.24.3

replace (
	k8s.io/api => ./staging/src/k8s.io/api
	k8s.io/apiextensions-apiserver => ./staging/src/k8s.io/apiextensions-apiserver
	k8s.io/apimachinery => ./staging/src/k8s.io/apimachinery
	k8s.io/apiserver => ./staging/src/k8s.io/apiserver
	k8s.io/cli-runtime => ./staging/src/k8s.io/cli-runtime
	k8s.io/client-go => ./staging/src/k8s.io/client-go
	k8s.io/cloud-provider => ./staging/src/k8s.io/cloud-provider
	k8s.io/cluster-bootstrap => ./staging/src/k8s.io/cluster-bootstrap
	k8s.io/code-generator => ./staging/src/k8s.io/code-generator
	k8s.io/component-base => ./staging/src/k8s.io/component-base
	k8s.io/component-helpers => ./staging/src/k8s.io/component-helpers
	k8s.io/controller-manager => ./staging/src/k8s.io/controller-manager
	k8s.io/cri-api => ./staging/src/k8s.io/cri-api
	k8s.io/cri-client => ./staging/src/k8s.io/cri-client
	k8s.io/csi-translation-lib => ./staging/src/k8s.io/csi-translation-lib
	k8s.io/dynamic-resource-allocation => ./staging/src/k8s.io/dynamic-resource-allocation
	k8s.io/endpointslice => ./staging/src/k8s.io/endpointslice
	k8s.io/externaljwt => ./staging/src/k8s.io/externaljwt
	k8s.io/kms => ./staging/src/k8s.io/kms
	k8s.io/kube-aggregator => ./staging/src/k8s.io/kube-aggregator
	k8s.io/kube-controller-manager => ./staging/src/k8s.io/kube-controller-manager
	k8s.io/kube-proxy => ./staging/src/k8s.io/kube-proxy
	k8s.io/kube-scheduler => ./staging/src/k8s.io/kube-scheduler
	k8s.io/kubectl => ./staging/src/k8s.io/kubectl
	k8s.io/kubelet => ./staging/src/k8s.io/kubelet
	k8s.io/metrics => ./staging/src/k8s.io/metrics
	k8s.io/mount-utils => ./staging/src/k8s.io/mount-utils
	k8s.io/pod-security-admission => ./staging/src/k8s.io/pod-security-admission
	k8s.io/sample-apiserver => ./staging/src/k8s.io/sample-apiserver
	k8s.io/sample-cli-plugin => ./staging/src/k8s.io/sample-cli-plugin
	k8s.io/sample-controller => ./staging/src/k8s.io/sample-controller
)

require (
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/client-go v0.0.0
	k8s.io/component-helpers v0.0.0
	k8s.io/klog/v2 v2.130.1
	k8s.io/utils v0.0.0-20250502105355-0f33e8f1c979
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/fxamacker/cbor/v2 v2.7.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/gnostic-models v0.6.9 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/oauth2 v0.27.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/term v0.30.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	golang.org/x/time v0.9.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/evanphx/json-patch.v4 v4.12.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/kube-openapi v0.0.0-20250318190949-c8a335a9a2ff // indirect
	sigs.k8s.io/json v0.0.0-20241010143419-9aa6b5e7a4b3 // indirect
	sigs.k8s.io/randfill v1.0.0 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.6.0 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
