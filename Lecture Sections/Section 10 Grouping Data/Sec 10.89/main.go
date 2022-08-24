package main

import (
	"fmt"
)

// Map - delete

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	delete(m, "James")
	fmt.Println(m)

	// delete something not in the map
	// no errors are thrown even tho this key isn't in the map
	delete(m, "Ian Fleming")
	fmt.Println(m)

	fmt.Println(m["Miss Moneypenny"])
	fmt.Println(m["Ian Fleming"])

	// using if statement to check if key is in map
	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("Value: ", v)
		delete(m, "Miss Moneypenny")
	}
	// check if key/value was deleted
	fmt.Println(m)
}
