package main

import (
	"fmt"
)

// Switch Statement Documentation

func main() {
	// a missing switch expression is equivalent value true
	switch {
	case false:
		fmt.Println("Should not print")
	case true:
		fmt.Println("Should print")
	}
}
