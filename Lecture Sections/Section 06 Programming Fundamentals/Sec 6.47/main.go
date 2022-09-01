package main

import "fmt"

// numeral systems

func main() {
	s := "H"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Println(n)
	//type
	fmt.Printf("%T\n", n)
	// binary
	fmt.Printf("%b\n", n)
	// hexidecimal
	fmt.Printf("%#X\n", n)
}
