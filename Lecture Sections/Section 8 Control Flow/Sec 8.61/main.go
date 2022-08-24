package main

import (
	"fmt"
)

// break and continue

func main() {
	x := 0
	for {
		// needs to increment from top for it to continue thru loop
		x++
		if x > 50 {
			break
		}
		// if not even number, do not print (check via modulus)
		if x%2 != 0 {
			continue
		}
		// will print even number
		fmt.Println(x)
	}
	fmt.Println("Done.")
}
