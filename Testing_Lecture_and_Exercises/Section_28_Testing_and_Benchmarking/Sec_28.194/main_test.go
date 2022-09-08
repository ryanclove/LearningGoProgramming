package main

import "testing"

// cmd: 	go test -v

// make first letter capital
// bets practice is to label it as the function you're testing
// aka test mySum function in main package
func TestMySum(t *testing.T) {
	if mySum(2, 3) != 5 {
		// https://pkg.go.dev/testing#T.Error
		// func (c *T) Error(args ...any)
		t.Error("Expected 2 + 3 = 5, got", mySum(2, 3))
	}
}

func TestFailureOfMySum(t *testing.T) {
	// obviously 2+3=5, but setting it to 6 so I see error
	if mySum(2, 3) != 6 {
		t.Error("Expected 2 + 3 = 6, got", mySum(2, 3))
		t.Error("I wanted this to fail to see the error print")
	}
}
