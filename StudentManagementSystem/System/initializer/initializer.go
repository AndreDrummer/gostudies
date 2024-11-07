package system_initializer

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	students "github.com/AndreDrummer/gostudies/StudentManagementSystem/Students"
	system_panel "github.com/AndreDrummer/gostudies/StudentManagementSystem/System"
	system "github.com/AndreDrummer/gostudies/StudentManagementSystem/System/structs"
	"github.com/AndreDrummer/gostudies/StudentManagementSystem/Utils/file_handler"
)

var (
	systemInstance *system.System
)

func initSystem() {
	systemInstance = system.NewSystem()
}

func getStudentNameAndGrades(studentInfo string) (string, string) {

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

func convertGradesToInt(grades string) []int {
	gradeStringSlice := strings.Fields(grades)
	gradeIntSlice := make([]int, 0)

	for _, v := range gradeStringSlice {
		gradeInt, err := strconv.Atoi(v)

		if err != nil {
			log.Fatal(err)
		}

		gradeIntSlice = append(gradeIntSlice, gradeInt)
	}

	return gradeIntSlice
}

func loadStudentsFromDB() {
	dbFile, err := os.OpenFile(system.DBFilename, os.O_RDONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	dbFileContent := file_handler.GetFileContent(dbFile)

	if len(dbFileContent) > 0 {
		for _, v := range dbFileContent {
			studentIDString := strings.Split(v, ".")[0]
			studentID, err := strconv.Atoi(studentIDString)
			studentName, grades := getStudentNameAndGrades(v)
			fmt.Println(studentName)
			fmt.Println(grades)

			if err != nil {
				log.Fatal(err)
			}

			newStudent := &students.Student{
				ID:     studentID,
				Grades: convertGradesToInt(grades),
				Name:   studentName,
			}

			systemInstance.StudentsQty++
			systemInstance.Students[studentID] = newStudent
		}
	}
}

func createDBFile(filename string) error {
	err := os.WriteFile(filename, []byte(""), 0644)

	if err != nil {
		fmt.Printf("ERROR creating file %v\n", filename)
		return err
	}

	return nil
}

// Fake DB: All is based on files
func initDB() {

	_, errorReadingFile := os.ReadFile(system.DBFilename)

	if errorReadingFile != nil {
		errorCreatingFile := createDBFile(system.DBFilename)

		if errorCreatingFile != nil {
			log.Fatal(errorCreatingFile)
		}

		errorReadingFile = nil
	}
}

func Initialize() {
	initSystem()

	initDB()
	loadStudentsFromDB()

	system_panel.Start(systemInstance)
}
