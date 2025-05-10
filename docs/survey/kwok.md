# [KWOK](https://kwok.sigs.k8s.io/)

## Setup

Install [kwok](https://kwok.sigs.k8s.io/docs/generated/kwok/) and [kwokctl](https://kwok.sigs.k8s.io/docs/generated/kwokctl/).

```zsh
% brew install kwok
% kwok --version
kwok version v0.6.1 go1.23.2 (darwin/arm64)
% kwokctl --version
kwokctl version v0.6.1 go1.23.2 (darwin/arm64)
```

## Usage

### Use kwokctl

```zsh
export KWOK_KUBE_VERSION=v1.33.0
kwokctl create cluster --name=kwok
kubectl config use-context kwok-kwok
kubectl cluster-info
kwokctl get clusters
kubectl version
kubectl get ns
kwokctl scale node --replicas=3
kubectl apply -f docs/survey/manifests/node.yaml
kubectl get nodes
kwokctl scale pod --replicas=3
kubectl apply -f docs/survey/manifests/pod.yaml
kubectl get pod -o wide
kwokctl delete cluster --name=kwok
```

### Use docker

```zsh
docker run --rm -it -p 8080:8080 ghcr.io/kwok-ci/cluster:v0.6.1-k8s.v1.33.0
```

```zsh
export KUBECONFIG=/path/to/docs/survey/kubeconfig.yaml  # Or use -s :8080 or --kubeconfig=docs/survey/kubeconfig.yaml
kubectl version
kubectl get ns
kubectl config view
```

## Scenario

### [Preemption](https://kwok.sigs.k8s.io/docs/technical-outcomes/scheduling/pod-priority-and-preemption/)

```zsh
kwokctl create cluster
kwokctl get clusters
kwokctl scale node --replicas 1 --param '.allocatable.cpu="4000m"'
kubectl get nodes
kubectl apply -f docs/survey/manifests/priority-classes.yaml
kubectl apply -f docs/survey/manifests/low-priority-pod.yaml
kubectl get pods
kubectl apply -f docs/survey/manifests/high-priority-pod.yaml
kubectl get pods
kubectl describe pod high-priority-pod | awk '/Events:/,/pod to node/'
kwokctl delete cluster
```

## How to

## Use specified scheduler

```zsh
kwokctl create cluster --kube-scheduler-binary=https://dl.k8s.io/v1.33.0/bin/linux/arm64/kube-scheduler
```
