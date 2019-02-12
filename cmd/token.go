// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"fmt"

	"github.com/spf13/cobra"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "return node bootstrap token",
	Long: `New node needs bootstrap token to register to Kubernetes cluster.
'token' command resolves and returns the token.`,
	RunE: func(_ *cobra.Command, args []string) error {
		token, err := getBootstrapToken()
		if err != nil {
			return err
		}

		fmt.Print(token)

		return nil
	},
	// We handle errors at root.go
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	configFlags.AddFlags(tokenCmd.Flags())
	addProvisionOptions(tokenCmd)

	rootCmd.AddCommand(tokenCmd)
}
