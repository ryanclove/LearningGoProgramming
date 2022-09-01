package main

import (
	"fmt"
)

// conditional if, else if, else

func main() {
	x := 42
	if x == 40 {
		fmt.Println("Our value was 40")
	} else if x == 41 {
		fmt.Println("Out value was 41")
	} else if x == 42 {
		fmt.Println("Out value was 42")
	} else if x == 43 {
		fmt.Println("Out value was 43")
	} else {
		fmt.Println("Out value was NOT 40, 41, 42, or 43")
	}
}
