package main

import (
	"fmt"
)

// Loop, conditional, modulus

func main() {
	for i := 0; i < 20; i++ {
		// divisible by 4
		if i%4 == 0 {
			fmt.Println(i)
		}
	}
}
