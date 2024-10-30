package simpleoperations

import (
	"fmt"

	"github.com/AndreDrummer/gostudies/SimpleOperations/utils"
)

func Calculator() {
	fmt.Println(" ** Welcome to GO! **")

	num1 := utils.GetIntInput("\nEnter First Number: ")
	num2 := utils.GetIntInput("\nEnter Second Number: ")

	fmt.Print("\n\nResults:\n")
	fmt.Printf("%d + %d = %d\n", num1, num2, (num1 + num2))
	fmt.Printf("%d - %d = %d\n", num1, num2, (num1 - num2))
	fmt.Printf("%d * %d = %d\n", num1, num2, (num1 * num2))
	if num2 == 0 {
		fmt.Println("Cannot perform division: divisor is zero!")
	} else {
		fmt.Printf("%d / %d = %d\n", num1, num2, (num1 / num2))
	}
}
