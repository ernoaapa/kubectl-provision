package provision

import (
	"errors"
	"testing"

	"github.com/ernoaapa/kubectl-provision/pkg/provision/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInstall(t *testing.T) {
	sampleErr := errors.New("sample error")

	tests := map[string]struct {
		err error
	}{
		"successful": {
			err: nil,
		},
		"with error": {
			err: sampleErr,
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		executorMock := &mocks.Executor{}
		i := Installer{
			executor: executorMock,
		}

		executorMock.On("Exec", mock.Anything).Return(test.err)

		err := i.Install()

		assert.Equal(t, test.err, err)

		executorMock.AssertExpectations(t)
	}
}
