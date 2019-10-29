package env

import "os"

// SetEnv sets environment variables
func SetEnv(addr, pemCert, token, role string) {
	os.Setenv("ADDR", addr)
	os.Setenv("CACERT", pemCert)
	os.Setenv("TOKEN", token)
	os.Setenv("ROLE", role)
}
