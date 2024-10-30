package simpleoperations

import (
	"fmt"

	"github.com/AndreDrummer/gostudies/SimpleOperations/utils"
)

func NumberType() {
	value := utils.GetIntInput("Enter a numeric value: ")

	positive := value > 0
	even := value%2 == 0

	if value == 0 {
		fmt.Printf("\n%d an even number!\n", value)
	} else {
		fmt.Printf("\n%d is an %v and %v number!\n", value,
			utils.Ternary(even, "even", "odd"),
			utils.Ternary(positive, "positive", "negative"),
		)
	}
}
