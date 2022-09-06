package main

import "fmt"

// Hands-on exercise #3
// at package level, assign values to the variables
// in func main, use fmt.Sprintf (String print) to print all values
// to one single string

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	// %v returns value of variable, \t is tab
	fmt.Println(s)

}
