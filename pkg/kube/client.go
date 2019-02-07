package kube

import (
	"time"

	"github.com/pkg/errors"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

type Client struct {
	config  *rest.Config
	timeout time.Duration
}

func NewClient(config *rest.Config) *Client {
	return &Client{
		config:  config,
		timeout: 60 * time.Second,
	}
}

func (c *Client) getClient(namespace string) (v1.SecretInterface, error) {
	clientset, err := kubernetes.NewForConfig(c.config)
	if err != nil {
		return nil, err
	}

	return clientset.CoreV1().Secrets(namespace), nil
}

// FindBootstrapToken find the bootstrap token from current cluster
func (c *Client) FindBootstrapToken() (apiv1.Secret, error) {
	client, err := c.getClient("kube-system")
	if err != nil {
		return apiv1.Secret{}, err
	}

	list, err := client.List(metav1.ListOptions{
		FieldSelector: "type=bootstrap.kubernetes.io/token",
	})
	if err != nil {
		return apiv1.Secret{}, err
	}

	switch len(list.Items) {
	case 0:
		return apiv1.Secret{}, errors.Wrap(ErrNotFound, "Bootstrap token Secret")
	case 1:
		return list.Items[0], nil
	default:
		return apiv1.Secret{}, ErrWithMessagef(ErrTooManyFound, "Bootstrap token Secrets (count %d)", len(list.Items))
	}
}
