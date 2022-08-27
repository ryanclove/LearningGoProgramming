package main

import (
	"fmt"
)

/*
Build and use an anonymous func
*/

func main() {
	func() {
		for i := 1; i <= 10; i++ {
			fmt.Println(i)
		}
	}()

	fmt.Println("Anonymous func ran")
}
