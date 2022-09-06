package main

import (
	"fmt"
)

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store
the following data:
	● first name
	● last name
	● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
which stores the favorite flavors.
*/

type person struct {
	first       string
	last        string
	favIceCream []string // slice of strings for multiple flavors
}

func main() {
	p1 := person{
		first: "ryan",
		last:  "coslove",
		favIceCream: []string{
			"Vanilla",
			"Cookies n' cream"},
	}
	p2 := person{
		first: "Go",
		last:  "Programmer",
		favIceCream: []string{
			"Chocolate",
			"Banana",
			"Pistachio"},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for _, v := range p1.favIceCream {
		fmt.Println(v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for _, v := range p2.favIceCream {
		fmt.Println(v)
	}

}
