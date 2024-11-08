package controller

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	utils "github.com/AndreDrummer/gostudies/StudentManagementSystem/Utils"
	"github.com/AndreDrummer/gostudies/StudentManagementSystem/Utils/file_handler"
	"github.com/AndreDrummer/gostudies/StudentManagementSystem/structs"
)

var inputRead *bufio.Reader = bufio.NewReader(os.Stdin)
var DBFilename = "StudentManagementSystem/System/db/students.txt"

type System struct {
	Students            map[int]*structs.Student
	StudentsQty         int
	MinimumPassingGrade int
}

func NewSystem() *System {
	return &System{
		Students:            make(map[int]*structs.Student),
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

func readYesOrNo(msg string) bool {
	fmt.Print(msg)
	asnwer, _ := inputRead.ReadString('\n')
	asnwer = strings.TrimSpace(asnwer)
	asnwerIsEmpty := asnwer == ""
	acceptableYesAnswers := []string{"YES", "Y", "yes", "y"}
	acceptableNOAnswers := []string{"NO", "N", "no", "n"}

	acceptableAnswers := []string{}
	acceptableAnswers = append(acceptableAnswers, acceptableNOAnswers...)
	acceptableAnswers = append(acceptableAnswers, acceptableYesAnswers...)

	if asnwerIsEmpty || !utils.Contains(acceptableAnswers, asnwer) {
		utils.ClearConsole()

		fmt.Println(" ** Invalid entry **, please try again.")
		for asnwerIsEmpty {
			fmt.Print("\nEnter student name: ")
			asnwer, _ = inputRead.ReadString('\n')
			asnwer = strings.TrimSpace(asnwer)
			asnwerIsEmpty = asnwer == ""
		}
	}

	if asnwer == "YES" || asnwer == "Y" || asnwer == "yes" || asnwer == "y" {
		return true
	} else if asnwer == "NO" || asnwer == "N" || asnwer == "no" || asnwer == "n" {
		return false
	}

	return false
}

func GetStudentNameAndGrades(studentInfo string) (string, string) {

	parts := strings.Fields(studentInfo)
	var gradeStartIndex int

	for i := 0; i < len(parts); i++ {
		if _, err := strconv.Atoi(parts[i]); err == nil {
			gradeStartIndex = i
			break
		}
	}

	var studentName, grades string

	if gradeStartIndex > 0 {
		studentName = strings.Join(parts[1:gradeStartIndex], " ")
		grades = strings.Join(parts[gradeStartIndex:], " ")
	} else {
		studentName = strings.Join(parts[1:], " ")
		grades = ""
	}

	return studentName, grades
}

func updateGradeOnDB(studentID int, gradePosition int, newGrade string) {
	dbFile := file_handler.OpenFileWithPerm(DBFilename, os.O_RDWR)

	if dbFile != nil {
		defer dbFile.Close()
		studentData := file_handler.GetFileEntryByPrefix(studentID, dbFile)
		studentName, grades := GetStudentNameAndGrades(studentData)
		updatedGrades := make([]string, 0)
		gradeSlice := strings.Fields(grades)

		if gradePosition < len(gradeSlice) {
			gradeSlice[gradePosition] = newGrade
		}

		updatedGrades = append(updatedGrades, gradeSlice...)
		updatedGrades = append(updatedGrades, newGrade)

		updatedGradesString := strings.Join(updatedGrades, " ")

		var buidler strings.Builder
		buidler.WriteString(fmt.Sprintf("%v. ", studentID))
		buidler.WriteString(studentName)
		buidler.WriteString(" ")
		buidler.WriteString(updatedGradesString)
		studentInfo := buidler.String()

		file_handler.UpdateFileEntry(dbFile, studentID, studentInfo)
	}
}

func removeStudentFromDB(studentID int) {
	dbFile := file_handler.OpenFileWithPerm(DBFilename, os.O_RDWR)

	if dbFile != nil {
		defer dbFile.Close()
		file_handler.RemoveFileEntry(dbFile, studentID)
	}
}

func getStudentByID(system *System, studentID int) (*structs.Student, bool) {
	student, exists := system.Students[studentID]
	return student, exists
}

func getNextAvailableID(system *System) int {
	studentIDs := make([]int, 0)

	for _, student := range system.Students {
		studentID := student.ID
		studentIDs = append(studentIDs, studentID)
	}

	sort.Ints(studentIDs)
	startID := 1
	for _, ID := range studentIDs {
		if ID-startID == 0 {
			startID++
			continue
		} else {
			return startID
		}
	}

	return len(studentIDs) + 1
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

	studentID := getNextAvailableID(system)

	newStudent := &structs.Student{
		ID:     studentID,
		Grades: make([]int, 0),
		Name:   studentName,
	}

	system.Students[newStudent.ID] = newStudent
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%d. ", newStudent.ID))
	builder.WriteString(fmt.Sprintf("%s ", newStudent.Name))

	dbFile := file_handler.OpenFileWithPerm(DBFilename, os.O_APPEND|os.O_WRONLY)

	if dbFile != nil {
		defer dbFile.Close()
		file_handler.AppendToFile(dbFile, builder.String())
	}

	utils.ClearConsole()
	utils.SetSuccessMsg(fmt.Sprintf("Student %v Added!", studentName))
	system.StudentsQty++
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

				gradeString := fmt.Sprintf("%d", grade)

				studentGradesQty := len(student.Grades)
				updateGradeOnDB(studentID, studentGradesQty, gradeString)

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
			removeStudentFromDB(studentID)
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

	if params == nil {
		params = &displayAllParams{}
	}

	if params.displayMsg == "" {
		msg = "\n\nPress Enter to go back. "
	} else {

		msg = params.displayMsg
	}

	if areThereStudentsRegistered(system) {
		tempSliceToSort := make([]string, 0)

		for _, v := range system.Students {
			var line string
			if len(v.Grades) == 0 {
				line = fmt.Sprintf("%d - %s -- No grades recorded.", v.ID, v.Name)
				tempSliceToSort = append(tempSliceToSort, line)
			} else {
				line = fmt.Sprintf("%d - %s --Grades-> %v", v.ID, v.Name, v.Grades)
				tempSliceToSort = append(tempSliceToSort, line)
			}

		}
		utils.SortSliceStringByID(tempSliceToSort, "-")

		for _, v := range tempSliceToSort {
			fmt.Println(v)
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

func (system *System) ClearDB() {
	answer := readYesOrNo("This will delete all data save in the database. Are you sure? ")
	if answer {
		system.Students = make(map[int]*structs.Student)
		DBFile := file_handler.OpenFileWithPerm(DBFilename, os.O_TRUNC)
		defer DBFile.Close()

		file_handler.ClearFileContent(DBFile)
		utils.SetSuccessMsg("\n** Operação realizada com sucesso! **")
	}
}
