package main

import (
	"fmt"
)

/*
A “callback” is when we pass a func into a func as an argument. For this exercise,
	● pass a func into a func as an argument
*/

func main() {
	s := func(s string) string {
		og := "This is the original string"
		return og
	}
	str := s("")
	fmt.Println(str)
	ns := newString(s, str)
	fmt.Println(ns)
}

func newString(f func(s string) string, og string) string {
	s1 := f(og)
	s2 := s1 + " + the concatenated string"
	return s2
}

/*
He had a different answer but same premise:

g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}

	x := foo(g, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)
}

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}

*/
