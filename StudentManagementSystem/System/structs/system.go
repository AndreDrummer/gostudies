package system

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	students "github.com/AndreDrummer/gostudies/StudentManagementSystem/Students"
	utils "github.com/AndreDrummer/gostudies/StudentManagementSystem/Utils"
	"github.com/AndreDrummer/gostudies/StudentManagementSystem/Utils/file_handler"
)

var inputRead *bufio.Reader = bufio.NewReader(os.Stdin)
var DBFilename = "StudentManagementSystem/students.txt"

type System struct {
	Students            map[int]*students.Student
	StudentsQty         int
	MinimumPassingGrade int
}

func NewSystem() *System {
	return &System{
		Students:            make(map[int]*students.Student),
		StudentsQty:         0,
		MinimumPassingGrade: 60,
	}
}

type displayAllParams struct {
	displayMsg string
	readInput  interface{}
}

func areThereStudentsRegistered(system *System) bool {
	return len(system.Students) > 0
}

func readStudentID(system *System) int {
	var studentID int

	for {
		system.DisplayAll(&displayAllParams{
			displayMsg: "\nEnter the student ID: ",
			readInput:  &studentID,
		})
		_, exists := getStudentByID(system, studentID)

		if exists {
			break
		}

		utils.ClearConsole()
		fmt.Print("\nPlease enter a valid student ID!\n\n")
	}

	return studentID
}

func readGrade() int {
	var grade int

	fmt.Print("Enter grade (0-100): ")
	_, err := fmt.Scanf("%d", &grade)

	for err != nil || grade < 0 || grade > 100 {
		utils.ClearConsole()
		fmt.Print("Invalid grade. Enter a number between 0 and 100: ")
		_, err = fmt.Scanf("%d", &grade)
	}

	return grade
}

func updateGrade(grades string, gradePosition int, newGrade string) string {
	gradeSlice := strings.Fields(grades)
	updatedGrades := make([]string, 0)

	gradeSlice[gradePosition] = newGrade
	updatedGrades = append(updatedGrades, gradeSlice...)

	return strings.Join(updatedGrades, " ")
}

func getStudentByID(system *System, studentID int) (*students.Student, bool) {
	student, exists := system.Students[studentID]
	return student, exists
}

func (system *System) AddStudent() {
	fmt.Print("\nEnter student name: ")
	studentName, _ := inputRead.ReadString('\n')
	studentName = strings.TrimSpace(studentName)
	nameIsEmpty := studentName == ""

	if nameIsEmpty {
		utils.ClearConsole()

		fmt.Println(" ** Invalid name **, please try again.")
		for nameIsEmpty {
			fmt.Print("\nEnter student name: ")
			studentName, _ = inputRead.ReadString('\n')
			studentName = strings.TrimSpace(studentName)
			nameIsEmpty = studentName == ""
		}
	}

	system.StudentsQty++
	newStudent := &students.Student{
		ID:     system.StudentsQty,
		Grades: make([]int, 0),
		Name:   studentName,
	}

	system.Students[newStudent.ID] = newStudent
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%d. ", newStudent.ID))
	builder.WriteString(fmt.Sprintf("%s ", newStudent.Name))

	dbFile, err := os.OpenFile(DBFilename, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("ERROR saving student to DB: %v", err)
	}
	defer dbFile.Close()
	file_handler.AppendToFile(dbFile, builder.String())

	utils.ClearConsole()
	utils.SetSuccessMsg(fmt.Sprintf("Student %v Added!", studentName))
}

func (system *System) AddGrade() {
	if areThereStudentsRegistered(system) {
		fmt.Print("What student would you like to add a grade?\n\n")
		studentID := readStudentID(system)
		student, studentExists := getStudentByID(system, studentID)

		if studentExists {
			grade := readGrade()

			if grade >= 0 {
				student.AddGrade(grade)

				utils.SetSuccessMsg(fmt.Sprintf("Grade %v added to %v!", grade, student.Name))
			} else {
				utils.ClearConsole()
			}
		}
	} else {
		utils.PressEnterToGoBack("\n** Empty! No student registered.")
	}
}

func (system *System) RemoveStudent() {
	if areThereStudentsRegistered(system) {
		studentID := readStudentID(system)
		student, studentExists := getStudentByID(system, studentID)

		if studentExists {
			delete(system.Students, studentID)
			utils.SetSuccessMsg(fmt.Sprintf("Student %v removed!", student.Name))
		}
	} else {
		utils.PressEnterToGoBack("\n** Empty! No student registered.")
	}
}

func (system *System) CalculateAverage() {
	if areThereStudentsRegistered(system) {
		studentID := readStudentID(system)
		student, studentExists := getStudentByID(system, studentID)

		if studentExists {
			avg := student.GetAverage()
			utils.PressEnterToGoBack(fmt.Sprintf("\nThe average of %s is %v.\n", student.Name, avg))
		}
	} else {
		utils.PressEnterToGoBack("\n** Empty! No student registered.")
	}
}

func (system *System) CheckPassOrFail() {

	if areThereStudentsRegistered(system) {
		studentID := readStudentID(system)
		student, studentExists := getStudentByID(system, studentID)

		if studentExists {
			passed := student.GetAverage() >= system.MinimumPassingGrade
			var resultMsg string
			if passed {
				resultMsg = "has been approved! :)"
			} else {
				resultMsg = "has failed :(!"
			}
			utils.PressEnterToGoBack(fmt.Sprintf("\n%s %v.\n", student.Name, resultMsg))
		}

	} else {
		utils.PressEnterToGoBack("\n** Empty! No student registered.")
	}
}

func (system *System) DisplayAll(params *displayAllParams) {
	var msg string

	if params != nil && params.displayMsg == "" {
		msg = "\n\nPress Enter to go back. "
	} else {
		params = &displayAllParams{}
		msg = params.displayMsg
	}

	if areThereStudentsRegistered(system) {
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
			utils.ClearConsole()
		} else {
			utils.PressEnterToGoBack("")
		}

	} else {
		utils.PressEnterToGoBack("\n** Empty! No student registered.")
	}
}
