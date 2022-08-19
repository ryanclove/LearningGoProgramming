package main

import "fmt"

func main() {
	a := (42 == 42) // true
	b := (42 <= 43) // true
	c := (42 >= 43) // false
	d := (42 != 43) // true
	e := (42 < 43)  // true
	f := (42 > 43)  // false
	fmt.Println(a, b, c, d, e, f)
}
