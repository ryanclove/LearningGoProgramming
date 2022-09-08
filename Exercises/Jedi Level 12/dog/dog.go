// Package dog provides functionality to allow us to understand human to dog years
package dog

import (
	"fmt"
)

// Years takes an inputted number of years (human) and converts it into dog years
func Years(n int) int {
	dogYears := n * 7

	fmt.Println("Converting...")

	fmt.Printf("Human years: %d\t Dog years: %d", n, dogYears)
	return dogYears
}
