// Copyright © 2018 ERNO AAPA <ERNO.AAPA@GMAIL.COM>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"

	"github.com/ernoaapa/kubectl-bootstrap/pkg/bootstrap"
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:   "node ADDRESS",
	Short: "Install and configures node for Kubernetes",
	Long:  `bootstrap node connects to existing server and bootstrap it for Kubernetes`,
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("ADDRESS is required for node bootstrap")
		}

		token, err := getBootstrapToken()
		if err != nil {
			return err
		}

		host, pins, err := getServer()
		if err != nil {
			return err
		}

		return bootstrap.NewInstaller(
			host,
			token,
			pins,
			bootstrap.NewSSHExecutor(args),
		).Install()
	},
	// We handle errors at root.go
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	configFlags.AddFlags(nodeCmd.Flags())
	addBootstrapOptions(nodeCmd)

	rootCmd.AddCommand(nodeCmd)
}
