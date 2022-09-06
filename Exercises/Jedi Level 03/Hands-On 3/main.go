package main

import (
	"fmt"
)

// Create a for loop using for condition {} syntax
// Print all the years I've been alive

func main() {
	bd := 2000
	for bd <= 2022 {
		fmt.Println(bd)
		bd++
	}
}
