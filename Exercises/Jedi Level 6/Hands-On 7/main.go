package main

import (
	"fmt"
)

/*
Assign a func to a variable, then call that func
*/

var g func() = func() {
	fmt.Println("g from outside") // this wont print when i make g = f
}

func main() {
	f := func() {
		for i := 1; i <= 3; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("%T\n", f)

	g = f
	g()
	fmt.Printf("%T\n", g)

}
