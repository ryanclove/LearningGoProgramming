package main

import (
	"fmt"
)

// Sort custom

type person struct {
	Name string
	Age  int
}

// Prints a string of a person's name and age
func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

func main() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
}
