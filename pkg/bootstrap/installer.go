package bootstrap

import (
	"fmt"
	"strings"
)

//go:generate mockery --name=Executor

// Executor is common interface for executors
type Executor interface {
	Exec([]string) error
}

// Installer install Kubernetes worker over executor
type Installer struct {
	host     string
	token    string
	certPins []string
	executor Executor
}

// NewInstaller return new Installer
func NewInstaller(host string, token string, certPins []string, executor Executor) *Installer {
	return &Installer{host, token, certPins, executor}
}

// Install executes all required commands for installing Kubernetes worker
func (i *Installer) Install() error {
	if err := i.installDockerRuntime(); err != nil {
		return err
	}

	if err := i.installKube(); err != nil {
		return err
	}

	if err := i.join(); err != nil {
		return err
	}

	return nil
}

func (i *Installer) installDockerRuntime() error {
	return i.executor.Exec([]string{
		"set -e",
		"sudo apt-get update && sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common",
		"curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -",
		"sudo add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\"",
		"sudo apt-get update && sudo apt-get install -y docker-ce=18.06.0~ce~3-0~ubuntu",
		"sudo sh -c \"echo '{\\\"exec-opts\\\": [\\\"native.cgroupdriver=systemd\\\"], \\\"log-driver\\\": \\\"json-file\\\", \\\"log-opts\\\": { \\\"max-size\\\": \\\"100m\\\" }, \\\"storage-driver\\\": \\\"overlay2\\\" }' > /etc/docker/daemon.json\"",

		"sudo mkdir -p /etc/systemd/system/docker.service.d",
		"sudo systemctl daemon-reload",
		"sudo systemctl restart docker",
	})
}

func (i *Installer) installKube() error {
	return i.executor.Exec([]string{
		"set -e",
		"sudo apt-get update && sudo apt-get install -y apt-transport-https curl",
		"curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -",
		"sudo sh -c 'echo \"deb https://apt.kubernetes.io/ kubernetes-xenial main\" > /etc/apt/sources.list.d/kubernetes.list'",
		"sudo apt-get update",
		"sudo apt-get install -y kubelet kubeadm kubectl",
		"sudo apt-mark hold kubelet kubeadm kubectl",
	})
}

func (i *Installer) join() error {
	certHashes := make([]string, 0, len(i.certPins))
	for _, pin := range i.certPins {
		certHashes = append(certHashes, fmt.Sprintf("--discovery-token-ca-cert-hash %s", pin))
	}
	return i.executor.Exec([]string{
		"set -x",
		fmt.Sprintf("sudo kubeadm join %s --token %s %s", i.host, i.token, strings.Join(certHashes, " ")),
	})
}
