package main

import (
	"fmt"
	"io"
	"os"
)

// Writer Interface

// writer interface works for all of these
// flexibility and accesibility in go lang

func main() {
	fmt.Println("Hello, from package fmt")
	fmt.Fprintln(os.Stdout, "Hello, from package os")
	io.WriteString(os.Stdout, "Hello, from package io")
}
