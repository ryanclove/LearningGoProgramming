package main

import (
	"fmt"
)

// There is no while function in Go

func main() {
	// for init; condition; post {}
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}
