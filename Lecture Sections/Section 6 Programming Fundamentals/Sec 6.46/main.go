package main

import (
	"fmt"
)

func main() {
	s := "Hello, string"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// Slice of bytes
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	/* prints:
	[72 101 108 108 111 44 32 115 116 114 105 110 103]
	[]uint8

	If we go to:
	https://en.wikipedia.org/wiki/ASCII

	In the table if we go to 72 (decimal) we find the letter "H"
	If we go to 101 we find letter "e"
	If we go to 108 we find letter "l"
	If we go to 108 we find letter "o"
	*/

	// Print UTF-8 Codepoint
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")

	// Print index Position i of string s and the hex value v
	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
