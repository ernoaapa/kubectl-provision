package bootstrap

import (
	"errors"
	"testing"

	"github.com/ernoaapa/kubectl-bootstrap/pkg/bootstrap/mocks"
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

	executorMock := &mocks.Executor{}
	i := Installer{
		executor: executorMock,
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		executorMock.On("Exec", mock.Anything).Return(test.err).Once()

		err := i.Install()

		assert.Equal(t, test.err, err)

		executorMock.AssertExpectations(t)
	}
}
