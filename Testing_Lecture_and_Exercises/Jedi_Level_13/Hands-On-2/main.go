package main

import (
	"fmt"

	"github.com/GoesToEleven/LearningGoProgramming/Testing_Lecture_and_Exercises/Jedi_Level_13/Hands-On-2/quote"
	"github.com/GoesToEleven/LearningGoProgramming/Testing_Lecture_and_Exercises/Jedi_Level_13/Hands-On-2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
