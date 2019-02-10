package cmd

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/client-go/rest"
)

var cert = []byte(`-----BEGIN CERTIFICATE-----
MIIC4DCCAcgCCQCdnRCHIuj1+zANBgkqhkiG9w0BAQsFADAyMQswCQYDVQQGEwJG
STERMA8GA1UEBwwISGVsc2lua2kxEDAOBgNVBAoMB1Rlc3RpbmcwHhcNMTkwMjEw
MDQyOTA5WhcNMjAwMjEwMDQyOTA5WjAyMQswCQYDVQQGEwJGSTERMA8GA1UEBwwI
SGVsc2lua2kxEDAOBgNVBAoMB1Rlc3RpbmcwggEiMA0GCSqGSIb3DQEBAQUAA4IB
DwAwggEKAoIBAQDoyzDDg3GP1LdzapsjobbB3Xr6CLLyNRZ1g3+eSwCG66Hsp/yL
Hrxxa2KN2WXNNlMC22d/tdE2wVKCfVSIyQa1DH/kzbCywEtEQFEqo8g2v1/77Fzf
ADZ8jFnVbHJVj5YM5hXHi/MjZCvjbG8ew7WTfRWRbOABC8sFEfqlyN3syC7g5Ea2
7obsXKrcsq7BB0Kex7G85KPQk8OPSxofpygAj20lrM7jlz89unJiKWZxrqBxP42w
BCBQ3/a/KKLyi+MM/N4Hu+aFJmLnXor7X98eIRJYrq/2SLqGsJv9jewuyIcYuixh
RXhKE3Dk04O/49Jwd/F5C8C5AMaRp5DNQNBhAgMBAAEwDQYJKoZIhvcNAQELBQAD
ggEBACI4EyxGE1LxST4HvnEPcBtVU7notPXVUIKgSidi/LAvRfyKXiocD7wnjL7z
O6Cz+bRPy0eimVaif6PJmjN6rJwuS/rXUMRIJvtFuYH08nppI4/1A7pjGl7+k4N4
iZuKkvMNpY9vR0HatvC+LRW1ivIQU+z2u0WXKz56bhkNwIZeqhEZn8IDffmZxrvo
gKPV/fTS+KjgdBiDV+5MrTNtfKD7fQN4x8x4adEsGavbBlHPvGorm1qGl5+XMLor
GvyV2Dut5vTJ2aEnY00g5GDrx7riZpJvqQXnCKQJA2xKO1kbmVj4u+iPVuYm8Y8c
TU5zVlHKIlI6Ci3DQPnjY9Jtgws=
-----END CERTIFICATE-----
`)
var certChecksum = "sha256:477108018e2325d2b8bf477fcbb190b9ed1d242cf14c44a3b790a2f316974f98"

func TestGetCACertHash(t *testing.T) {
	caFile := writeToTempFile(t, cert)
	defer os.Remove(caFile)

	tests := map[string]struct {
		config   *rest.Config
		expected []string
	}{
		"success data field": {
			config: &rest.Config{
				TLSClientConfig: rest.TLSClientConfig{
					CAData: []byte(base64.StdEncoding.EncodeToString(cert)),
				},
			},
			expected: []string{certChecksum},
		},
		"success file field": {
			config: &rest.Config{
				TLSClientConfig: rest.TLSClientConfig{
					CAFile: caFile,
				},
			},
			expected: []string{certChecksum},
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)

		pins, err := getCACertPins(test.config)
		require.NoError(t, err)
		assert.Equal(t, test.expected, pins)
	}
}

func writeToTempFile(t *testing.T, content []byte) string {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "test-")
	if err != nil {
		t.Fatal(err)
	}
	if _, err = tmpFile.Write(content); err != nil {
		t.Fatal(err)
	}

	return tmpFile.Name()
}
