package main

import (
	"fmt"
)

// we create VALUES of a certain TYPE that are stored in VARIABLES
// and those VARIABLES have IDENTIFIERS

// Refer to README for going over some housekeeping links under documentation section

/*
In Go:
1. you don't create classes, you create a TYPE
2. you don't instantiate, you create a VALUE of a TYPE
*/

var x int

type person struct {
	first string
	last  string
	age   int
}

// not good practice, alias-ing a type
type foo int

// not good practice, alias-ing a type
var y foo

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Println(p1)
}
