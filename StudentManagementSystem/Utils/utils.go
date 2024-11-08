package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

func ClearConsole() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		fmt.Println("It was not possible to clear the console on this OS kkkkkk.")
	}

	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to clear console:", err)
	}

}

func PressEnterToGoBack(msg string) {
	fmt.Print(msg)
	fmt.Print("\n\nPress Enter to go back. ")
	fmt.Scanln()
	ClearConsole()
}

func ShowOperationResultMsg(format string, a ...interface{}) {
	ClearConsole()
	fmt.Printf(format, a...)
	time.Sleep(1 * time.Second)
	ClearConsole()
}

func SetSuccessMsg(msg string) {
	ShowOperationResultMsg("** Success ** %v!", msg)
}

func SortSliceStringByID(slice []string, splitter string) {
	sort.Slice(slice, func(i, j int) bool {
		ID1 := strings.Split(slice[i], splitter)[0]
		ID2 := strings.Split(slice[j], splitter)[0]

		ID1 = strings.TrimSpace(ID1)
		ID2 = strings.TrimSpace(ID2)

		ID1NUM, err1 := strconv.Atoi(ID1)
		ID2NUM, err2 := strconv.Atoi(ID2)

		if err1 != nil || err2 != nil {
			fmt.Printf("Error converting IDs to integers: %v, %v\n", err1, err2)
			return false
		}

		return ID1NUM < ID2NUM
	})
}

func Contains(slice []string, something string) bool {
	for _, v := range slice {
		if v == something {
			return true
		}
	}

	return false
}
