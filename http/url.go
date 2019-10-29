package http

// URL returns a correctly formatted url for Vault http requests based on conffiguration and internal path
func URL(address, path string) string {
	version := "/v1/"

	return address + version + path
}
