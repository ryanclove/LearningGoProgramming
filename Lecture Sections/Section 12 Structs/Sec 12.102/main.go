package main

import (
	"fmt"
)

// documentation: https://go.dev/ref/spec#Struct_types

type person struct {
	// field names must be unique aka first, second, age are all unique
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{ //initialize fields with values
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2.first, p2.last, p2.age)
}
