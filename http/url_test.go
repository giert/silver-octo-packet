package http

import "testing"

func Test_MakeUrl(t *testing.T) {
	// Test empty config
	addr := ""
	path := ""
	want := "/v1/"
	got := URL(addr, path)
	if got != want {
		t.Fatalf("wanted %s, got %s", want, got)
	}

	/**
	// Test good config
	addr = Mock.Addr
	path = Mock.Path
	want = addr + "/v1/" + path
	got = URL(addr, path)
	if got != want {
		t.Fatalf("wanted %s, got %s", want, got)
	}
	*/
}
