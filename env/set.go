package env

import "os"

// SetEnv sets environment variables for creating a Vault client
func SetEnv(vaultAddr, pemCert, githubToken, k8ServicePath, k8MountPath, k8Role string) {
	os.Setenv("VAULT_ADDR", vaultAddr)
	os.Setenv("VAULT_CACERT", pemCert)
	os.Setenv("GITHUB_TOKEN", githubToken)
	os.Setenv("SERVICE_ACCOUNT_PATH", k8ServicePath)
	os.Setenv("MOUNT_PATH", k8MountPath)
	os.Setenv("ROLE", k8Role)
}
