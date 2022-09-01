package main

import (
	"fmt"
)

// printing ascii

func main() {
	// https://en.wikipedia.org/wiki/ASCII#:~:text=041-,33,-21
	for i := 33; i <= 122; i++ {
		// %v prints value
		// %#x prints hexidecimal value
		// %#U prints unicode ascii value
		fmt.Printf("%v\t %#x\t %#U\n", i, i, i)
		// ascii wise, look at wiki link
		// will print punctuation and alphabet
	}
}
