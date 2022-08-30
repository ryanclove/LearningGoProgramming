package main

import (
	"encoding/json"
	"fmt"
)

// JSON Marshal
// func Marshal(v any) ([]byte, error)

// Upercase identifiers are exported (i.e, First, Last, Age)
// lowercase are NOT exported data (i.e, first, last, age)
type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}

	// Marshal is going to return a slice of bytes and an error
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
