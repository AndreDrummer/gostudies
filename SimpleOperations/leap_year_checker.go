package simpleoperations

import (
	"fmt"
	"time"

	"github.com/AndreDrummer/gostudies/SimpleOperations/utils"
)

func isLeapYear(year int) bool {
	return (year%400 == 0) || (year%4 == 0 && year%100 != 0)
}

func LeapYearChecker() {
	currentYear := time.Now().Year()
	year := utils.GetIntInput(fmt.Sprintf("Enter a valid year. Ex.: %d: ", currentYear))

	if isLeapYear(year) {
		fmt.Printf("%d %v", year, "is a leap year.\n")
	} else {
		fmt.Printf("%d %v", year, " is NOT a leap year.\n")
	}
}
