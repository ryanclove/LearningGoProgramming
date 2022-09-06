package main

import (
	"fmt"
)

/*
Create and use an anonymous struct.
*/

func main() {

	s := struct {
		first       string
		friends     map[string]int
		favIceCream []string
	}{
		first: "Ryan",
		friends: map[string]int{
			"Matt S":   22,
			"Matt A":   23,
			"Analisse": 22,
		},
		favIceCream: []string{
			"Cookie Dough",
			"Strawberry",
		},
	}

	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favIceCream)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for _, val := range s.favIceCream {
		fmt.Println(val)
	}

}
