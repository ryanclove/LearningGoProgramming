package main

import "testing"

/*
Table Tests

We can write a series of tests to run. This allows us to test a variety of situations
*/

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		// tests that will pass
		test{[]int{21, 21}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{1, 1}, 2},
		test{[]int{-5, 7, 13, -4}, 11},
		test{[]int{-1, 0, 1}, 0},
		// purposefully fail this test in the table
		//test{[]int{-1, 0, 1}, 2},
	}

	for _, v := range tests {
		x := mySum(v.data...) // unfurl each of the ints from the slice
		if x != v.answer {
			t.Error("Expected:", v.answer, "\tGot", x)
		}
	}
}
