package main

import (
	"fmt"
)

// composite or aggregate data structure
type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age) // create values of type person
	fmt.Println(p2.first, p2.last, p2.age) // NOT called objects
}
