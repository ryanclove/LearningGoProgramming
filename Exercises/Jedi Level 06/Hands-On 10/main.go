package main

import (
	"fmt"
)

/*
Closure is when we have “enclosed” the scope of a variable in some code block. For this
hands-on exercise, create a func which “encloses” the scope of a variable
*/

func main() {
	// code block level scope
	{
		y := 42
		fmt.Println(y)
	}

	// return incrementer() value which returns func() int in its scope
	b := incrementor()
	c := incrementor()
	fmt.Println(b(), c())
	fmt.Println(b(), c())
	fmt.Println(b())
	fmt.Println(b(), c())
}

func incrementor() func() int {
	var a int
	return func() int {
		a++
		return a
	}
}
