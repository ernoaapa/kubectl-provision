package bootstrap

//go:generate mockery --name=Executor

type Executor interface {
	Exec([]string) error
}

type Installer struct {
	executor Executor
}

func (i *Installer) Install() error {
	return i.executor.Exec([]string{
		"sudo apt-get update && sudo apt-get install -y apt-transport-https curl",
		"curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -",
		"sudo sh -c 'echo \"deb https://apt.kubernetes.io/ kubernetes-xenial main\" > /etc/apt/sources.list.d/kubernetes.list'",
		"sudo apt-get update",
		"sudo apt-get install -y kubelet kubeadm kubectl",
		"sudo apt-mark hold kubelet kubeadm kubectl",
	})
}
