package cmd

import (
	"github.com/ernoaapa/kubectl-bootstrap/pkg/cmd"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
)

type bootstrapOptions struct {
	BootstrapToken string
}

var configFlags = genericclioptions.NewConfigFlags(false)
var opt = bootstrapOptions{}

func addBootstrapOptions(cmd *cobra.Command) {
	cmd.Flags().StringVar(&opt.BootstrapToken, "bootstrap-token", opt.BootstrapToken, "The bootstrap token to use, rather than resolving from Kubernetes cluster.")
}

func getKubeConfig() (*rest.Config, error) {
	config, err := configFlags.ToRESTConfig()
	if err != nil {
		return nil, err
	}
	rest.SetKubernetesDefaults(config)
	return config, nil
}

func getBootstrapToken() (string, error) {
	if opt.BootstrapToken != "" {
		return opt.BootstrapToken, nil
	}

	config, err := getKubeConfig()
	if err != nil {
		return "", err
	}

	return cmd.GetBootstrapToken(config)
}
