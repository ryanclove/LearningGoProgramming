package main

import (
	"fmt"
)

// anonymous structs
// if you only need a struct for 1 little area, it prevents
// code pollution or have extraneous types where you don't need them

func main() {
	// put what person is representing in the place of person struct
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	// it's anonymous because it does not have a name (previously person)
	fmt.Println(p1)
}
