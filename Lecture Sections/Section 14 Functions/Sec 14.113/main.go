package main

import (
	"fmt"
)

// Methods

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) indentifier(parameters) (return(s)) { code }

// any value of this type secretAgent has access to this method
func (s secretAgent) speak() {
	fmt.Println("I am secret agent", s.first, s.last)
}

func main() {

	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
