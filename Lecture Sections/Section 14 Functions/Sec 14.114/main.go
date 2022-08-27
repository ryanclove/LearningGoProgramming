package main

import (
	"fmt"
)

// Interfaces & polymorphism

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am secret agent", s.first, s.last)
}

// new function
func (p person) speak() {
	fmt.Println("I am a person named", p.first, p.last)
}

// ANY type that has the method speak() is also of type human
// a value can be more than one type
type human interface {
	speak()
	// values of type secretAgent are also of type human
	// values of type person now are also type of human as well
}

func bar(h human) {
	// switch statement where you can switch on type
	// insertion
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar because my type, person, is human ~", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar because my type, secretAgent, is human ~", h.(secretAgent).first, h.(secretAgent).last)
	}
	fmt.Println("I was passed into bar", h)
}

// used below for conversion
type hotdog int

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

	// new person
	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	// print Dr. Yes
	fmt.Println(p1)

	// pass secret agent and person into bar function that are humans
	bar(sa1)
	bar(sa2)
	bar(p1)

	// conversion - hotdog to int
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
