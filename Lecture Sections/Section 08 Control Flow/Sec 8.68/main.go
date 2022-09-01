package main

import (
	"fmt"
)

// Conditional logic operators

func main() {
	fmt.Printf("true && true\t %v\n", true && true) // and both conditions
	fmt.Printf("true && false\t %v\n", true && false)
	fmt.Printf("true || true\t %v\n", true || true) // or either condition is true
	fmt.Printf("true || false\t %v\n", true || false)
	fmt.Printf("!true\t\t %v\n", !true)
	fmt.Printf("!false\t\t %v\n", !false)
}
