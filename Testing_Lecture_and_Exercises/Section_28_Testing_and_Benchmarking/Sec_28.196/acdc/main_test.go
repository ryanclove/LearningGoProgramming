package acdc

import (
	"fmt"
)

/*
Example Tests

Examples show up in documentation
	● godoc -http :8080
	● https://blog.golang.org/examples
	● go test ./…
		- test everything where i'm at./ and everything
		  down the directory tree
*/

func ExampleSum() {
	fmt.Println(Sum(2, 3))
	// Output:
	// 5
}

func ExampleSumWhereFails() {
	fmt.Println(Sum(2, 3))
	// Output:
	// 7
}
