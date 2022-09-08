package main

import "testing"

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
