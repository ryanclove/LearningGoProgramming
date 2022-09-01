package main

import (
	"fmt"
	// GOOS and GOARCH Packages
	"runtime"
)

// Numeric types

// can specify types, but not necessary for execution here for x and y

var x int
var y float64

// if I want exact precision in bits (for memory/storage purposes)
// signed 8-bit goes from -128 to 127
var z int8 = -128

func main() {
	x := 42
	y := 42.34534
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	// type int
	fmt.Printf("%T\n", y)
	// type float64
	//x = 42.55
	// throws an error because now I've declared
	// the type above as int, and now x is not an int
	//fmt.Printf("%T\n", x)
	int68()
	runtimePackage()
}

func int68() {
	fmt.Println("--------------")
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}

func runtimePackage() {
	fmt.Println("--------------")
	fmt.Println(runtime.GOOS)
	// prints OS running the program
	// for me, outputs 'windows'
	fmt.Println(runtime.GOARCH)
	// runs the computer architecture running the program
	// for me, outputs 'amd64' processor (64 bit)
	// if was amd64p32, for example, would mean all pointers are 32-bit
}
