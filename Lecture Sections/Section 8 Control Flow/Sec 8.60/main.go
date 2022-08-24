package main

import (
	"fmt"
)

// for statement

func main() {
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("Done.")

	fmt.Println("Method 2 with if-statement: ")
	// Method 2: with if-statement
	x = 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("Done.")
}

/* The Go for loop is similar to—but not the same as—C's.
It unifies for and while and there is no do-while.
There are three forms, only one of which has semicolons.

// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }

*/
