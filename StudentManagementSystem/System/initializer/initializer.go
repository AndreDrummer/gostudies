package initializer

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/AndreDrummer/gostudies/StudentManagementSystem/Utils/file_handler"
	"github.com/AndreDrummer/gostudies/StudentManagementSystem/structs"
	system_panel "github.com/AndreDrummer/gostudies/StudentManagementSystem/system"
	student_system_controller "github.com/AndreDrummer/gostudies/StudentManagementSystem/system/controller"
)

var (
	systemInstance *student_system_controller.System
)

func initSystem() {
	systemInstance = student_system_controller.NewSystem()
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
	dbFile, err := os.OpenFile(student_system_controller.DBFilename, os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}

	dbFileContent := file_handler.GetFileContent(dbFile)

	// Remove any empty line that may exists.
	file_handler.OverrideFileContent(dbFile, dbFileContent)

	if len(dbFileContent) > 0 {
		for _, v := range dbFileContent {
			studentIDString := strings.Split(v, ".")[0]
			studentID, err := strconv.Atoi(studentIDString)
			if v == "" {
				continue
			}
			studentName, grades := student_system_controller.GetStudentNameAndGrades(v)

			if err != nil {
				log.Fatal(err)
			}

			newStudent := &structs.Student{
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

	_, errorReadingFile := os.ReadFile(student_system_controller.DBFilename)

	if errorReadingFile != nil {
		errorCreatingFile := createDBFile(student_system_controller.DBFilename)

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
