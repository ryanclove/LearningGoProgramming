package main

import (
	"fmt"
)

// Print every rune code point of the uppercase alphabet 3 times
// https://en.wikipedia.org/wiki/ASCII#:~:text=101-,65,-41

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
