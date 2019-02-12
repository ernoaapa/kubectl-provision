package provision

import (
	"os"
	"os/exec"
	"strings"
)

type SSH struct {
	args []string
}

func NewSSHExecutor(args []string) *SSH {
	return &SSH{args}
}

func (s *SSH) Exec(cmds []string) error {
	command := exec.Command("ssh", s.args...)
	command.Stdin = strings.NewReader(strings.Join(cmds, ";"))
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
