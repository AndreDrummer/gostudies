package system

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	students "github.com/AndreDrummer/gostudies/StudentManagementSystem/Students"
)

type System struct {
	Students            map[int]*students.Student
	studentsQty         int
	minimumPassingGrade int
}

type displayAllParams struct {
	displayMsg string
	readInput  interface{}
}

var (
	inputRead      *bufio.Reader
	systemInstance *System
)

func StartStudentManagementSystem() {
	initSystem()
	inputRead = bufio.NewReader(os.Stdin)
	displayOptions()
}

func initSystem() *System {
	if systemInstance == nil {
		systemInstance = &System{
			Students:            make(map[int]*students.Student),
			studentsQty:         0,
			minimumPassingGrade: 60,
		}
	}

	return systemInstance
}

func showOperationResultMsg(format string, a ...interface{}) {
	clearConsole()
	fmt.Printf(format, a...)
	time.Sleep(1 * time.Second)
	clearConsole()
}

func setSuccessMsg(msg string) {
	showOperationResultMsg("** Success ** %v!", msg)
}

func pressEnterToGoBack(msg string) {
	fmt.Print(msg)
	fmt.Print("\n\nPress Enter to go back. ")
	fmt.Scanln()
	clearConsole()
}

func displayOptions() {
	clearConsole()

	fmt.Print("Welcome to the Student Management System!\n")

	for {
		fmt.Print("\nChoose an option below: \n\n")
		fmt.Println("1 - Add a new Student")
		fmt.Println("2 - Add a grade to a Student")
		fmt.Println("3 - Remove a Student")
		fmt.Println("4 - Calculate average score of a Student")
		fmt.Println("5 - Check if a student passed or failed")
		fmt.Println("6 - Display all students and their grades")
		fmt.Println("0 - Exit")
		fmt.Print("\nEnter your choice: ")

		var choice int
		fmt.Scanln(&choice)
		handleChoice(choice)
	}
}

func handleChoice(choice int) {
	clearConsole()
	switch choice {
	case 0:
		fmt.Printf("\n\n ** Goodbye! **\n\n")
		time.Sleep(750 * time.Millisecond)
		os.Exit(0)
	case 1:
		systemInstance.addStudent()
	case 2:
		systemInstance.addGrade()
	case 3:
		systemInstance.removeStudent()
	case 4:
		systemInstance.calculateAverage()
	case 5:
		systemInstance.checkPassOrFail()
	case 6:
		systemInstance.displayAll(&displayAllParams{})
	default:
		fmt.Println("Invalid choice. Try again.")
	}
}

func clearConsole() {
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

func areThereStudentsRegistered() bool {
	return len(systemInstance.Students) > 0
}

func readStudentID() int {
	var studentID int

	for {
		systemInstance.displayAll(&displayAllParams{
			displayMsg: "\nEnter the student ID: ",
			readInput:  &studentID,
		})
		_, exists := systemInstance.getStudentByID(studentID)

		if exists {
			break
		}

		clearConsole()
		fmt.Print("\nPlease enter a valid student ID!\n\n")
	}

	return studentID
}

func readGrade() int {
	var grade int

	fmt.Print("Enter grade (0-100): ")
	_, err := fmt.Scanf("%d", &grade)

	for err != nil || grade < 0 || grade > 100 {
		clearConsole()
		fmt.Print("Invalid grade. Enter a number between 0 and 100: ")
		_, err = fmt.Scanf("%d", &grade)
	}

	return grade
}

func (system *System) getStudentByID(studentID int) (*students.Student, bool) {
	student, exists := system.Students[studentID]
	return student, exists
}

func (system *System) addStudent() {
	fmt.Print("\nEnter student name: ")
	studentName, _ := inputRead.ReadString('\n')
	studentName = strings.TrimSpace(studentName)
	nameIsEmpty := studentName == ""

	if nameIsEmpty {
		clearConsole()

		fmt.Println(" ** Invalid name **, please try again.")
		for nameIsEmpty {
			fmt.Print("\nEnter student name: ")
			studentName, _ = inputRead.ReadString('\n')
			studentName = strings.TrimSpace(studentName)
			nameIsEmpty = studentName == ""
		}
	}

	system.studentsQty++
	newStudent := &students.Student{
		ID:     system.studentsQty,
		Grades: make([]int, 0),
		Name:   studentName,
	}

	system.Students[newStudent.ID] = newStudent

	clearConsole()
	setSuccessMsg(fmt.Sprintf("Student %v Added!", studentName))
}

func (system *System) addGrade() {
	if areThereStudentsRegistered() {
		fmt.Print("What student would you like to add a grade?\n\n")
		studentID := readStudentID()
		student, studentExists := system.getStudentByID(studentID)

		if studentExists {
			grade := readGrade()

			if grade >= 0 {
				student.AddGrade(grade)

				setSuccessMsg(fmt.Sprintf("Grade %v added to %v!", grade, student.Name))
			} else {
				clearConsole()
			}
		}
	} else {
		pressEnterToGoBack("\n** Empty! No student registered.")
	}
}

func (system *System) removeStudent() {
	if areThereStudentsRegistered() {
		studentID := readStudentID()
		student, studentExists := system.getStudentByID(studentID)

		if studentExists {
			delete(system.Students, studentID)
			setSuccessMsg(fmt.Sprintf("Student %v removed!", student.Name))
		}
	} else {
		pressEnterToGoBack("\n** Empty! No student registered.")
	}
}

func (system *System) calculateAverage() {
	if areThereStudentsRegistered() {
		studentID := readStudentID()
		student, studentExists := system.getStudentByID(studentID)

		if studentExists {
			avg := student.GetAverage()
			pressEnterToGoBack(fmt.Sprintf("\nThe average of %s is %v.\n", student.Name, avg))
		}
	} else {
		pressEnterToGoBack("\n** Empty! No student registered.")
	}
}

func (system *System) checkPassOrFail() {

	if areThereStudentsRegistered() {
		studentID := readStudentID()
		student, studentExists := system.getStudentByID(studentID)

		if studentExists {
			passed := student.GetAverage() >= system.minimumPassingGrade
			var resultMsg string
			if passed {
				resultMsg = "has been approved! :)"
			} else {
				resultMsg = "has failed :(!"
			}
			pressEnterToGoBack(fmt.Sprintf("\n%s %v.\n", student.Name, resultMsg))
		}

	} else {
		pressEnterToGoBack("\n** Empty! No student registered.")
	}
}

func (system *System) displayAll(params *displayAllParams) {
	var msg string

	if params.displayMsg == "" {
		msg = "\n\nPress Enter to go back. "
	} else {
		msg = params.displayMsg
	}

	if areThereStudentsRegistered() {
		for _, v := range system.Students {
			if len(v.Grades) == 0 {
				line := fmt.Sprintf("%d - %s -- No grades recorded.", v.ID, v.Name)
				fmt.Println(line)
			} else {
				line := fmt.Sprintf("%d - %s --Grades-> %v", v.ID, v.Name, v.Grades)
				fmt.Println(line)
			}

		}

		if params.readInput != nil {
			fmt.Print(msg)
			fmt.Scanln(params.readInput)
			clearConsole()
		} else {
			pressEnterToGoBack("")
		}

	} else {
		pressEnterToGoBack("\n** Empty! No student registered.")
	}
}
