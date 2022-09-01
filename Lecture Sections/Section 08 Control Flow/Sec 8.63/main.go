package main

import (
	"fmt"
)

// conditional if-statements

func main() {
	booleanChecks()
	initializationStatements()
}

// boolean conditions
func booleanChecks() {
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	// not true aka false
	if !true {
		fmt.Println("003")
	}
	// not false aka true
	if !false {
		fmt.Println("004")
	}
	if 2 == 2 {
		fmt.Println("005")
	}
	if 2 != 2 {
		fmt.Println("006")
	}
	if !(2 == 2) {
		fmt.Println("007")
	}
	if !(2 != 2) {
		fmt.Println("008")
	}
}

func initializationStatements() {
	// decalre an initial variable x := 42
	if x := 42; x == 42 {
		fmt.Println("009")
	}
	// fmt.Println(x)
}
