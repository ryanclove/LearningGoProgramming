package main

import (
	"fmt"
)

/*
Take the code from the previous exercise, then store the values of type person
in a map with the key of last name. Access each value in the map. Print
out the values, ranging over the slice.
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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for _, v := range p2.favIceCream {
			fmt.Println(v)
		}
	}

}
