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
	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/spf13/cobra"
)

type runOptions struct {
}

var configFlags = genericclioptions.NewConfigFlags(false)
var opt = runOptions{}

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Install and configures node for Kubernetes",
	Long:  `bootstrap node connects to existing server and bootstrap it for Kubernetes`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	// We handle errors at root.go
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	configFlags.AddFlags(nodeCmd.Flags())
	rootCmd.AddCommand(nodeCmd)
}
