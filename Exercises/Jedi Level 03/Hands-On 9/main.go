package main

import (
	"fmt"
)

// Create a program that uses a switch statement with the switch expression specified
// as a variable of TYPE string with the IDENTIFIER "favSport"

func main() {
	favSport := "Volleyball"
	switch favSport {
	case "Football":
		fmt.Println("Favorite Sport is Football")
	case "Basketball":
		fmt.Println("Favorite Sport is Basketball")
	case "Baseball":
		fmt.Println("Favorite Sport is Baseball")
	case "Hockey":
		fmt.Println("Favorite Sport is Hockey")
	case "Volleyball":
		fmt.Println("Favorite Sport is Volleyball")
	default:
		fmt.Printf("Favorite Sport is %s", favSport)
	}
}
