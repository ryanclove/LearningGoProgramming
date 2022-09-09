// Package acdc will sum a slice of ints
package acdc

// Sum will add the integers inside a slice of ints
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
