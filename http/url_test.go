package http

import (
	"testing"
	//mock "github.com/giert/silver-octo-packet/mock"
)

func Test_MakeUrl(t *testing.T) {
	// Test empty config
	addr := ""
	path := ""
	want := "/v1/"
	got := URL(addr, path)
	if got != want {
		t.Fatalf("wanted %s, got %s", want, got)
	}

	// Test good config
	//addr = mock.Addr
	//path = mock.Path
	want = addr + "/v1/" + path
	got = URL(addr, path)
	if got != want {
		t.Fatalf("wanted %s, got %s", want, got)
	}
}
