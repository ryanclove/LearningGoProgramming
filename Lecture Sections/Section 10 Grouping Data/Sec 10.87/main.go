package main

import (
	"fmt"
)

// Map - intro
// key values store
// it is an unordered list

func main() {
	m := map[string]int{
		//key          // value
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m) // returns the map

	fmt.Println(m["James"]) // returns 32

	fmt.Println(m["Barnabas"]) // returns 0 because that key/value is not stored in the map

	// check if value is stored in the map
	// comma, okay ", ok" idiom
	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok) // false

	if v, ok := m["Barnabas"]; ok {
		fmt.Println("This is the IF print but won't print", v)
	}

	if v, ok := m["James"]; ok {
		fmt.Println("This is the IF print and will print", v)
	}
}
