package bootstrap

func Node(args []string, token string) error {
	i := Installer{
		executor: NewSSHExecutor(args),
	}

	return i.Install()
}
