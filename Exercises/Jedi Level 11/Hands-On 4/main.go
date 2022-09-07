package main

import (
	"fmt"
	"log"
)

/*
Starting with this code - https://go.dev/play/p/JWxNVddnWqG

use the sqrt.Error struct as a value of type error. If you would like, use
these numbers for your
	● lat "50.2289 N"
	● long "99.4656 W"

I took out the lat and long strings from struct
*/

type sqrtError struct {
	err error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v", se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here

		// either way is fine

		// e := errors.New("Square root error")
		// return 0, sqrtError{e}

		sqrte := fmt.Errorf("square root of negative number: %v", f)
		return 0, sqrtError{sqrte}
	}
	return 42, nil
}
