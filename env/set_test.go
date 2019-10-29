package env

import (
	"os"
	"testing"

	"github.com/giert/silver-octo-packet/assert"
	"github.com/giert/silver-octo-packet/mock"
)

func Test_SetEnv(t *testing.T) {
	tests := []struct {
		addr    string
		pemCert string
		token   string
		role    string
	}{
		{
			addr:    mock.Addr,
			pemCert: mock.File,
			token:   mock.Token,
			role:    mock.Role,
		}, {
			addr:    "",
			pemCert: "",
			token:   "",
			role:    "",
		},
	}

	for _, test := range tests {
		SetEnv(test.addr, test.pemCert, test.token, test.role)
		assert.Result(t, os.Getenv("ADDR"), test.addr)
		assert.Result(t, os.Getenv("CACERT"), test.pemCert)
		assert.Result(t, os.Getenv("TOKEN"), test.token)
		assert.Result(t, os.Getenv("ROLE"), test.role)
	}
}
