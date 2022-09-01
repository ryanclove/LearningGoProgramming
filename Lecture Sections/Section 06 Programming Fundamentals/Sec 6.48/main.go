package main

import "fmt"

// Constants

//untyped constants
const (
	a = 42
	b = 42.78
	c = "James Bond"
)

//typed constants
const (
	d int     = 42
	e float64 = 42.78
	f string  = "James Bond"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	typed()
}

func typed() {
	fmt.Println("")
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
}
