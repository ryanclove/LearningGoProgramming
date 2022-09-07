package main

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

// getting this error in terminal:
// malformed import path "github.com/GoesToEleven/LearningGoProgramming/Exercises/Jedi Level 11/Hands-On 5": invalid char ' '
// think the files having a space in their names is screwing it up
