package main

import "fmt"

// Bit shifting

const (
	_ = iota
	k = 1 << (iota * 10) // kilobytes
	m = 1 << (iota * 10) // megaytes
	g = 1 << (iota * 10) // gigabytes
)

func main() {
	x := 4
	fmt.Printf("Decimal: %d\t\t Binary: %b\n", x, x)

	y := x << 1
	fmt.Printf("Decimal: %d\t\t Binary: %b\n", y, y)

	memoryBytes()
	// using iota for bit shifting
	iotaMemoryBytes()
}

func memoryBytes() {
	// kilobytes is 1024 bytes
	kb := 1024
	// megabytes * 1024 kilobytes
	mb := 1024 * kb
	// gigabytes * 1024 megabytes
	gb := 1024 * mb

	fmt.Printf("Decimal: %d\t\t Binary: %b\n", kb, kb)
	fmt.Printf("Decimal: %d\t Binary: %b\n", mb, mb)
	fmt.Printf("Decimal: %d\t Binary: %b\n", gb, gb)
	// each binary is shifting over by 10 places
}

func iotaMemoryBytes() {
	fmt.Println("Using iota: ")
	fmt.Printf("Decimal: %d\t\t Binary: %b\n", k, k)
	fmt.Printf("Decimal: %d\t Binary: %b\n", m, m)
	fmt.Printf("Decimal: %d\t Binary: %b\n", g, g)
}
