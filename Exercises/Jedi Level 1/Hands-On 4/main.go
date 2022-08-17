package main

import "fmt"

// Hands-on exercise #4

type banana int

// int is the underlying type of my new created type banana

var x banana

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
