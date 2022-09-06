package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("Decimal: %d \t Binary: %b \t Hexidecimal: %#x \n", a, a, a)
	b := a << 1
	fmt.Printf("Decimal: %d \t Binary: %b \t Hexidecimal: %#x \n", b, b, b)
}
