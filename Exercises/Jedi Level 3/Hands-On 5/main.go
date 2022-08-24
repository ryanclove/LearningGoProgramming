package main

import (
	"fmt"
)

// Print out remainder (modulus) between 10 and 100 when divded by 4

func main() {
	for i := 10; i <= 100; i++ {
		remainder := i % 4
		fmt.Printf("When %v is divided by 4, the remainder is (aka modulus) is %v\n", i, remainder)
	}
}
