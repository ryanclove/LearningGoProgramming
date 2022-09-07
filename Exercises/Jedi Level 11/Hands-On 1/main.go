package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
Start with this code - https://go.dev/play/p/6Z5kzW33bEP

Instead of using the blank identifier, make sure the code is checking and
handling the error
*/

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("JSON did not marshal ", err)
	}
	fmt.Println(string(bs))

}
