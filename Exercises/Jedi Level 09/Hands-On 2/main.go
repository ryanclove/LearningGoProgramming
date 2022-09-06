package main

import (
	"fmt"
)

/*
This exercise will reinforce our understanding of method sets:
	● create a type person struct
	● attach a method speak to type person using a pointer receiver
		○ *person
	● create a type human interface
		○ to implicitly implement the interface, a human must have the speak method
	● create func “saySomething”
		○ have it take in a human as a parameter
		○ have it call the speak method
	● show the following in your code
		○ you CAN pass a value of type *person into saySomething
		○ you CANNOT pass a value of type person into saySomething
	● here is a hint if you need some help
		○ https://play.golang.org/p/FAwcQbNtMG

Receivers Values
-----------------------------------------------
(t T) T and *T
(t *T) *T

*/

type person struct {
	First string
	Says  []string
}

func (p *person) speak() {
	// printing the first (index 0) saying in slice of strings
	fmt.Println(p.Says[0])
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		First: "James",
		Says:  []string{"Hello, I'm James Bond"},
	}
	saySomething(&p1)

	// does not work
	// saySomething(p1)

	p1.speak()
}
