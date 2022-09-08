package main

import (
	"fmt"

	"github.com/GoesToEleven/LearningGoProgramming/Testing_Lecture_and_Exercises/Section_28_Testing_and_Benchmarking/Sec_28.196/acdc"
)

/*
Example Tests

Examples show up in documentation
	● godoc -http :8080
	● https://blog.golang.org/examples
	● go test ./…
		- test everything where i'm at./ and everything
		  down the directory tree
*/

func main() {

	fmt.Println("2 + 3 =", acdc.Sum(2, 3))
	fmt.Println(acdc.Sum(5, 9, 11, 25))

}
