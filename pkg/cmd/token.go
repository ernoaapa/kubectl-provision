package cmd

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/ernoaapa/kubectl-bootstrap/pkg/kube"
	"k8s.io/client-go/rest"
)

// GetBootstrapToken resolve from current cluster the bootstrap token
func GetBootstrapToken(config *rest.Config) (string, error) {
	c := kube.NewClient(config)

	secrets, err := c.FindBootstrapTokens()
	if err != nil {
		return "", err
	}

	switch len(secrets) {
	case 1:
		secret := secrets[0]
		tokenID := string(secret.Data["token-id"])
		tokenSecret := string(secret.Data["token-secret"])
		return fmt.Sprintf("%s.%s", tokenID, tokenSecret), nil
	case 0:
		return "", errors.Wrap(ErrNotFound, "Bootstrap token Secrets")
	default:
		return "", ErrWithMessagef(ErrTooManyFound, "Bootstrap token Secrets (count %d)", len(secrets))
	}
}
