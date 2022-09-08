package main

import (
	"dogPackage/dog"
	"fmt"
)

// Main function
// Asks for human years input from user
func main() {
	var humanYear int

	fmt.Println("Convert human years to dog years:\nWhat age (human) do you want to convert?")

	_, err := fmt.Scanf("%d", &humanYear)

	if err != nil {
		fmt.Println(err)
	}

	// Call Package Dog from Main and execute Years function
	dog.Years(humanYear)
}
