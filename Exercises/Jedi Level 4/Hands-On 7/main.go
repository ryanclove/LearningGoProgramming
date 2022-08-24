package main

import (
	"fmt"
)

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
slice:
	"James", "Bond", "Shaken, not stirred"
	"Miss", "Moneypenny", "Helloooooo, James."
Range over the records, then range over the data in each record.
*/

func main() {
	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(x1)
	fmt.Println(x2)
	x := [][]string{x1, x2}
	fmt.Println(x)

	for i, xs := range x {
		fmt.Println("Record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: %v \n", j, val)
		}
	}
}
