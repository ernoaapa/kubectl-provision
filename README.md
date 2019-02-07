kubectl (Kubernetes CLI) plugin which bootstraps node.

It creates temporary _Pod_ and synchronises your local files to the desired container and executes any command.

### Why
It's best practice to control servers with configuration management tool, but sometimes you just want to connect existing server to Kubernetes cluster easily. `kubectl bootstrap node` install required components, get bootstrap token and connects the node to the target cluster.

### Other similar
#### [kubeadm](https://telepresence.io)

## How it works
`kubectl bootstrap node` connects to the target server over SSH and uses local package manager to install `kubelet` and configures the `kubelet` to bootstrap itself with bootstrap token.

## Install

### With Krew (Kubernetes plugin manager)
```shell
krew update
krew install bootstrap
```

### MacOS with Brew
```shell
brew install rsync ernoaapa/kubectl-plugins/bootstrap
```
### Linux / MacOS without Brew
1. Install rsync with your preferred package manager
2. Download `kubectl-bootstrap` binary from [releases](https://github.com/ernoaapa/kubectl-bootstrap/releases)
3. Add it to your `PATH`

## Usage
When the plugin binary is found from `PATH` you can just execute it through `kubectl` CLI
```shell
kubectl bootstrap node --help
```

## Development
### Prerequisites
- Golang v1.11
- [Go mod enabled](https://github.com/golang/go/wiki/Modules)

### Build and run locally
```shell
go run ./main.go node -h

# Syncs your local files to Kubernetes and list the files
```

### Build and install locally
```shell
go install .

# Now you can use `kubectl`
kubectl bootstrap --help
```
