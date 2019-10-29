package http

import (
	"testing"

	"github.com/giert/silver-octo-packet/assert"
	"github.com/giert/silver-octo-packet/mock"
)

func Test_MakeUrl(t *testing.T) {
	// Test empty
	addr := ""
	path := ""
	want := "/v1/"
	got := URL(addr, path)
	assert.Result(t, got, want)

	// Test mock data
	addr = mock.Addr
	path = mock.Path
	want = addr + "/v1/" + path
	got = URL(addr, path)
	assert.Result(t, got, want)
}
