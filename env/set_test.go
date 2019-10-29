package env

import (
	"testing"
)

func Test_SetEnv(t *testing.T) {
	/**
	tests := []struct {
		vaultAddr   string
		pemCert     string
		githubToken string
	}{
		{
			vaultAddr:   mock.Addr,
			pemCert:     mockFile,
			githubToken: mockToken,
		}, {
			vaultAddr:   "",
			pemCert:     "",
			githubToken: "",
		},
	}

	for _, test := range tests {
		SetEnv(test.vaultAddr, test.pemCert, test.githubToken, "", "", "")

		if os.Getenv("VAULT_ADDR") != test.vaultAddr || os.Getenv("VAULT_CACERT") != test.pemCert || os.Getenv("GITHUB_TOKEN") != test.githubToken {
			t.Fatalf("Unexpected environment variable")
		}
	}
	*/
}
