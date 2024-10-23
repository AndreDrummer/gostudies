package system

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	students "github.com/AndreDrummer/gostudies/StudentManagementSystem/Students"
)

type System struct {
	BD          map[int]students.Student
	studentsQty int
}

var operationStatus string = ""
var operationMsg string = ""
var systemIntance *System
var inputRead *bufio.Reader

func initSystem() *System {
	if systemIntance == nil {
		systemIntance = &System{
			BD:          make(map[int]students.Student),
			studentsQty: 0,
		}
	}

	return systemIntance

}

func showOperationResultMsg() {
	clearConsole()
	fmt.Printf(" %v %v ", operationStatus, operationMsg)
	operationStatus = ""
	operationMsg = ""
	time.Sleep(1 * time.Second)
	clearConsole()
}

func StartStudentManagementSystem() {
	initSystem()
	inputRead = bufio.NewReader(os.Stdin)
	displayOptions()
}

func displayOptions() {
	clearConsole()
	var op int = -1
	fmt.Print("Welcome to the Student Management System!\n")
	for op != 0 {

		fmt.Print("\nChoose an option below: \n\n")
		fmt.Println("1 - Add a new Student")
		fmt.Println("2 - Add a grade to a Student")
		fmt.Println("3 - Remove a Student")
		fmt.Println("4 - Calculate average score of a Student")
		fmt.Println("5 - Check if a student passed or failed")
		fmt.Println("6 - Display all students and their grades")
		fmt.Println("0 - Exit")
		fmt.Print("\nEnter your choice: ")

		op, err := fmt.Scanf("%d", &op)

		if err != nil {
			panic(err)
		} else {
			handleChoice(op)
		}
	}
}

func handleChoice(choice int) {
	clearConsole()
	switch choice {
	case 0:
		os.Exit(0)
	case 1:
		systemIntance.addStudent()
	case 2:
		systemIntance.addGrade()
	case 3:
		systemIntance.removestudent()
	case 4:
		systemIntance.calculateAverage()
	case 5:
		systemIntance.checkPassOrFail()
	case 6:
		systemIntance.displayAll()
	}
}

// func getNewStudentID() int {
// 	mapKeys := make([]int, 0)

// 	for key := range systemIntance.BD {
// 		mapKeys = append(mapKeys, key)
// 	}

// 	return len(mapKeys)
// }

func clearConsole() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "/cls")
	default:
		fmt.Println("It was not possible to clear the console on this OS.")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (system *System) addStudent() {

	fmt.Print("Enter student name: ")
	studentName, _ := inputRead.ReadString('\n')

	clearConsole()

	newStudent := &students.Student{
		ID:     system.studentsQty + 1,
		Grades: make([]int, 0),
		Name:   studentName,
	}

	system.BD[newStudent.ID] = *newStudent
	system.studentsQty = system.studentsQty + 1
	operationStatus = "** Success **"
	operationMsg = "Student Added!"
	showOperationResultMsg()
}

func (system *System) addGrade() {
	var studentID int
	var grade int
	fmt.Println("What student would you like to add a grade?")
	system.displayAll()
	fmt.Print("\nEnter the student ID ")
	studentID, err := fmt.Scanf("%d", &studentID)

	if err != nil {
		panic(err)
	}

	fmt.Print("What is the grade? ")
	grade, err = fmt.Scanf("%d", &grade)

	if err != nil {
		panic(err)
	}

	student := system.BD[studentID]
	studentGrades := student.Grades
	studentGrades = append(studentGrades, grade)
	student.Grades = studentGrades
}

func (system *System) removestudent() {

}

func (system *System) calculateAverage() {

}

func (system *System) checkPassOrFail() {

}

func (system *System) displayAll() {
	if len(system.BD) > 0 {
		for _, v := range system.BD {
			fmt.Printf("\nStudentID %v: - Student Name: %v\n", v.ID, v.Name)
		}
	} else {
		fmt.Print("** Empty! No student filed.")

	}

	fmt.Print("\n\nPress Enter to go back. ")
	fmt.Scanln()
	clearConsole()
}
