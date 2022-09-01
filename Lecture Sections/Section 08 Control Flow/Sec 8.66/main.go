package main

import (
	"fmt"
)

// Conditional - switch statement

func main() {
	fmt.Println("Main method: ")
	switch {
	case false:
		fmt.Println("This should not print b/c it's false")
	case (2 == 4):
		fmt.Println("This should not print b/c 2 != 4")
	case (3 == 3):
		fmt.Println("Prints, 3 == 3 is true")
	case (4 == 4):
		fmt.Println("4 == 4 is also true, does it print? no because first true case did not specify a fallthrough")
	}
	fmt.Println("-----------------")
	fmt.Println("fallthroughASwitchCase method: ")
	fallthroughASwitchCase()
	fmt.Println("-----------------")
	fmt.Println("stringCases method: ")
	stringCases()
	fmt.Println("-----------------")
	fmt.Println("multipleCases method: ")
	multipleCases()
}

// word to the wise, don't use fallthrough
func fallthroughASwitchCase() {
	switch {
	case (5 == 5):
		fmt.Println("Prints, 5 == 5 is true, and specifies a fallthrough")
		fallthrough
	case (6 == 8):
		fmt.Println("This prints even though 6 != 8, b/c of the fallthorugh from the previous case")
		fallthrough
	case (7 == 7):
		fmt.Println("Prints, 7 == 7 is true, and previous case specifies a fallthrough")
		fallthrough
	default:
		fmt.Println("this is the default")
	}
}

func stringCases() {
	n := "Bond"
	switch n {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("bond james")
	case "Q":
		fmt.Println("This is q")
	default:
		fmt.Println("This is default from String Cases")
	}
}

func multipleCases() {
	n := "Bond"
	switch n {
	case "Moneypenny", "Bond", "Dr No":
		fmt.Println("Moneypenny or Bond or Dr No")
	case "M":
		fmt.Println("This is m")
	case "Q":
		fmt.Println("This is q")
	default:
		fmt.Println("This is default from String Cases")
	}
}
