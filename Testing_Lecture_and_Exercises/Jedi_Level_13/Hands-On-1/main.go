package starting_code

import (
	"fmt"

	"github.com/GoesToEleven/LearningGoProgramming/Testing_Lecture_and_Exercises/Jedi_Level_13/Hands-On-1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
