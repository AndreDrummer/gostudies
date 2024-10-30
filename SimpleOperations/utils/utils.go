package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Ternary(condition bool, doThis any, doThat any) any {
	if condition {
		return doThis
	} else {
		return doThat
	}
}

func GetIntInput(prompt string) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)

		if err == nil {
			return num
		}

		fmt.Print("\n** Invalid input! **\n")
	}
}

func GetRuneInput(prompt string, specifics ...rune) rune {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		isSpecificChar := len(specifics) == 0

		if input == "" {
			fmt.Println("\n** Input cannot be empty! **")
			continue
		}

		char := rune(input[0])

		if len(specifics) > 0 {
			for _, v := range specifics {
				if v == char {
					isSpecificChar = true
					break
				} else {
					isSpecificChar = false
					break
				}
			}
		}

		if err == nil && isSpecificChar {
			return char
		}

		fmt.Print("\n** Invalid input! **\n")
	}
}
