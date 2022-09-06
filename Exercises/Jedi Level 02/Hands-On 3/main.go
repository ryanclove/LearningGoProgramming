package main

import "fmt"

//untyped constants
const (
	a     = 42 // untyped
	b int = 42 //typed
)

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

}
