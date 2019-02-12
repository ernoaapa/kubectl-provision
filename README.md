# `kubectl provision` (experimental)
Experimental `kubectl` (Kubernetes CLI) plugin which provisions an node with `kubeadm` from zero.

## Why
It's best practice to manage server configurations with some configuration management tool (Ansible, etc.), but sometimes you just want to connect an existing server to Kubernetes cluster quickly and easily.
`kubeadm` do the heavy job for joining the node to the cluster, but before that, you need to install container runtime, kubelet and kubeadm.

### How it works
 `kubectl provision node` will:
1. Install [required packages](https://kubernetes.io/docs/setup/independent/install-kubeadm/) (runtime, kubelet, kubeadm)
2. Get Kubernetes [bootstrap token](https://kubernetes.io/docs/reference/access-authn-authz/bootstrap-tokens/)
3. Joins the node to the target cluster with [kubeadm](https://kubernetes.io/docs/setup/independent/kubelet-integration/#workflow-when-using-kubeadm-join).

## Install

### MacOS with Brew
```shell
brew install rsync ernoaapa/kubectl-plugins/provision
```
### Linux / MacOS without Brew
1. Install rsync with your preferred package manager
2. Download `kubectl-provision` binary from [releases](https://github.com/ernoaapa/kubectl-provision/releases)
3. Add it to your `PATH`

## Usage
When the plugin binary is found from `PATH` you can just execute it through `kubectl` CLI
```shell
kubectl provision node --help
```

## Development
### Prerequisites
- Golang v1.11
- [Go mod enabled](https://github.com/golang/go/wiki/Modules)

### Build and run against Vagrant
#### Prerequisites
- [Vagrant](https://vagrantup.com)
- Existing Kubernetes cluster (e.g. [minikube](https://kubernetes.io/docs/setup/minikube/)

You need to have following flags in your Kubernetes master to be able to join with bootstrap tokens
```shell
# kube-apiserver
--enable-bootstrap-token-auth=true

# kube-controller-manager
--controllers=*,bootstrapsigner,tokencleaner
```

Start the test node
```shell
vagrant up
```

Install and join the Vagrant VM to your Kubernetes cluster 
```shell
go run ./main.go node -- -F =(vagrant ssh-config) node-1
```
