package main

import (
	"fmt"
	"math"
)

/*
● create a type SQUARE
● create a type CIRCLE
● attach a method to each that calculates AREA and returns it
	○ circle area= π r 2
	○ square area = L * W
● create a type SHAPE that defines an interface as anything that has the AREA method
● create a func INFO which takes type shape and then prints the area
create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle
● create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle
*/

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (sq square) area() float64 {
	squarea := sq.width * sq.length
	return squarea
}

func (ci circle) area() float64 {
	circarea := math.Pi * math.Pow(ci.radius, 2)
	return circarea
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c := circle{
		radius: 6.4,
	}
	s := square{
		length: 4.25,
		width:  4.25,
	}
	info(c)
	info(s)
}
