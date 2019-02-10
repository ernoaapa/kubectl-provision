package cmd

import (
	"crypto/x509"
	"encoding/base64"
	"net/url"

	"github.com/ernoaapa/kubectl-bootstrap/pkg/cmd"
	"github.com/kubernetes/kubernetes/cmd/kubeadm/app/util/pubkeypin"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
	clientcertutil "k8s.io/client-go/util/cert"
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

func getServer() (string, []string, error) {
	config, err := getKubeConfig()
	if err != nil {
		return "", []string{}, err
	}

	address, err := url.Parse(config.Host)
	if err != nil {
		return "", []string{}, err
	}

	pins, err := getCACertPins(config)
	if err != nil {
		return config.Host, []string{}, err
	}

	return address.Host, pins, nil
}

func getCACertPins(config *rest.Config) ([]string, error) {
	// load CA certificates from the kubeconfig (either from PEM data or by file path)
	var (
		caCerts []*x509.Certificate
		err     error
	)
	if config.CAData != nil {
		ca, err := base64.StdEncoding.DecodeString(string(config.CAData))
		if err != nil {
			return []string{}, errors.Wrap(err, "Failed to decode base64 certificate from kubeconfig")
		}
		caCerts, err = clientcertutil.ParseCertsPEM(ca)
		if err != nil {
			return []string{}, errors.Wrap(err, "failed to parse CA certificate from kubeconfig")
		}
	} else if config.CAFile != "" {
		caCerts, err = clientcertutil.CertsFromFile(config.CAFile)
		if err != nil {
			return []string{}, errors.Wrap(err, "failed to load CA certificate referenced by kubeconfig")
		}
	} else {
		return []string{}, errors.New("no CA certificates found in kubeconfig")
	}

	// hash all the CA certs and include their public key pins as trusted values
	publicKeyPins := make([]string, 0, len(caCerts))
	for _, caCert := range caCerts {
		publicKeyPins = append(publicKeyPins, pubkeypin.Hash(caCert))
	}
	return publicKeyPins, nil
}
