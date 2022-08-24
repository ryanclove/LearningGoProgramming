package main

import (
	"fmt"
)

// Create a for loop using for {} syntax aka while
// Print all the years I've been alive

func main() {
	bd := 2000
	for {
		if bd > 2022 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
