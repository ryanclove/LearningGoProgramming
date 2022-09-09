package main

import (
	"fmt"

	"github.com/GoesToEleven/LearningGoProgramming/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.198/saying"
)

/*
Benchmark

Part of the testing package allows us to measure the speed of our code. This could also be
called “measuring the performance” of your code, or “benchmarking” your code - finding out how
fast the code runs.

Runs the code a lot of times until it gets a statistcally accurate
measurement of how long it took to run it - give it a loop

go test -bench .

*/

func main() {

	fmt.Println(saying.Greet("James"))

}
