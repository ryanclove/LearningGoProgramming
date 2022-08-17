package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

// copied from the golang blog on using go-modules.

// to run this and the test, use 'go mod init' in git bash
// to initialize a go.mod file
// ex.: go mod init example.com/hello
// Note that it does not need to be a domain, could be example/repo

// then run 'go test' in terminal

// to remove a module, use 'rm -rf go.mod'
