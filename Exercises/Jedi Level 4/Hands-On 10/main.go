package main

import (
	"fmt"
)

// Using the code from the previous example, add a record to your map. Now print the map out
// using the “range” loop

func main() {
	m := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	// fmt.Println(m)

	// Add a record
	m["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	// delete a record (deleting dr.no)
	delete(m, "no_dr")

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
