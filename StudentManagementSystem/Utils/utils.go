package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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
		fmt.Println("It was not possible to clear the console on this OS.")
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
