package main

import (
	"fmt"
	"math"
)

// Method sets
/*
The method set of a type determines the methods that can be called on an operand of that type. Every type has a (possibly empty) method set associated with it:

The method set of a defined type T consists of all methods declared with receiver type T.
The method set of a pointer to a defined type T (where T is neither a pointer nor an interface) is the set of all methods declared with receiver *T or T.
The method set of an interface type is the intersection of the method sets of each type in the interface's type set (the resulting method set is usually just the set of declared methods in the interface).
Further rules apply to structs (and pointer to structs) containing embedded fields, as described in the section on struct types. Any other type has an empty method set.

In a method set, each method must have a unique non-blank method name.

Receivers Values
-----------------------------------------------
(t T) T and *T
(t *T) *T

*/

type circle struct {
	radius float64
}

type shape interface {
	areaNoRNoPV() float64
	areaNoPRbutPV() float64
	areaWithPRAndPV() float64
	areaWithPRButNoPV() float64
}

// ------------------------------------------
// Non Pointer Receiver & Non-Pointer Value
func (ci circle) areaNoRNoPV() float64 {
	circarea := math.Pi * math.Pow(ci.radius, 2)
	return circarea
}

func nonPointerReceiverAndNonPointerValue() {
	c := circle{
		radius: 6.4,
	}
	fmt.Printf("Type: %T\n", c)
	fmt.Println("Area:", c.areaNoRNoPV(), "\n")
}

// ------------------------------------------
// Non Pointer Receiver & Pointer Value
func (ci circle) areaNoPRbutPV() float64 {
	circarea := math.Pi * math.Pow(ci.radius, 2)
	return circarea
}

func nonPointerReceiverAndPointerValue() {
	c := circle{
		radius: 6.4,
	}
	fmt.Printf("Type: %T\n", &c)
	fmt.Println("Area:", c.areaNoPRbutPV(), "\n")

}

// ------------------------------------------
// Pointer Receiver & Pointer Value

// pointer of a *circle
func (ci *circle) areaWithPRAndPV() float64 {
	circarea := math.Pi * math.Pow(ci.radius, 2)
	return circarea
}

func pointerReceiverAndPointerValue() {
	c := circle{
		radius: 6.4,
	}
	fmt.Printf("Type: %T\n", &c)
	fmt.Println("Area:", c.areaWithPRAndPV(), "\n")
}

// ------------------------------------------
// Pointer Receiver & Pointer Value

// pointer of a *circle
func (ci *circle) areaWithPRButNoPV() float64 {
	circarea := math.Pi * math.Pow(ci.radius, 2)
	return circarea
}

func pointerReceiverButNoPointerValue() {
	c := circle{
		radius: 6.4,
	}
	fmt.Printf("Type: %T\n", c)
	fmt.Println("Area:", c.areaWithPRButNoPV(), "\n")
}

// ------------------------------------------
// Func main to run all 4 func combos of Pointer Receiver and Pointer Values
func main() {
	fmt.Println("NON-POINTER RECEIVER & NON-POINTER VALUE (c circle, info of c) :")
	nonPointerReceiverAndNonPointerValue()
	fmt.Println("NON-POINTER RECEIVER & POINTER VALUE (c circle, info of &c) :")
	nonPointerReceiverAndPointerValue()
	fmt.Println("POINTER RECEIVER & POINTER VALUE (c *circle, info of &c) :")
	pointerReceiverAndPointerValue()
	fmt.Println("POINTER RECEIVER & NON-POINTER VALUE (c *circle, info of c) :")
	pointerReceiverButNoPointerValue()
}
