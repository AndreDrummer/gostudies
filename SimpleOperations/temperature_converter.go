package simpleoperations

import (
	"fmt"

	"github.com/AndreDrummer/gostudies/SimpleOperations/utils"
)

func celsiusToFahrenheit(celsiusDegree int) int {
	return (celsiusDegree * 9 / 5) + 32
}
func fahrenheitToCelsius(fahrenheitDegree int) int {
	return (fahrenheitDegree - 32) * 5 / 9
}

func TemperatureConverter() {
	temp := utils.GetIntInput("Enter the temperature value: ")
	unit := utils.GetRuneInput("Convert from (C/F): ", 'c', 'f', 'C', 'F')

	if unit == 'c' || unit == 'C' {
		fmt.Printf("\n%dºC is %dºF\n", temp, celsiusToFahrenheit(temp))
	} else if unit == 'f' || unit == 'F' {
		fmt.Printf("\n%dºF is %dºC\n", temp, fahrenheitToCelsius(temp))
	} else {
		fmt.Printf("\nUnknown measure!\n")
	}
}
