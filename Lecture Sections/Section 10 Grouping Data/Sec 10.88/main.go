package main

import (
	"fmt"
)

// Map - add element and range

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	// add element
	m["todd"] = 33

	if v, ok := m["Barnabas"]; ok {
		fmt.Println("This is the IF print but won't print b/c not ok", v)
	}

	// print keys and values using for range
	for k, v := range m {
		fmt.Printf("The key %s returns the value %v\n", k, v)
	}

	// Printing a slice using for range
	xi := []int{4, 5, 7, 8, 9, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}
