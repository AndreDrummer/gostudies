package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetFileContent(file *os.File) []string {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	var content []string

	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, line)
	}

	return content
}

func PrintFileContent(file *os.File) {
	fileContent := GetFileContent(file)

	for _, v := range fileContent {
		fmt.Println(v)
	}
}

func GetFileContentUpdated(file *os.File, studentID int, updatedcontent string) []string {
	fileContent := GetFileContent(file)
	var newContent []string

	for _, v := range fileContent {
		if strings.HasPrefix(v, fmt.Sprintf("%d.", studentID)) {
			newContent = append(newContent, updatedcontent)
		} else {
			newContent = append(newContent, v)
		}
	}

	return newContent
}

func getStudentByID(ID int, file *os.File) string {
	fileContent := GetFileContent(file)

	for _, v := range fileContent {
		if strings.HasPrefix(v, fmt.Sprintf("%d.", ID)) {
			return v
		}
	}

	return ""
}

func getStudentGrades(studentInfo string) (string, string) {

	parts := strings.Fields(studentInfo)
	var gradeStartIndex int

	for i := 0; i < len(parts); i++ {
		if _, err := strconv.Atoi(parts[i]); err == nil {
			gradeStartIndex = i
			break
		}
	}
	student := strings.Join(parts[:gradeStartIndex], " ")
	grades := strings.Join(parts[gradeStartIndex:], " ")

	return student, grades
}

func updateGrade(grades string, gradePosition int, newGrade string) string {
	gradeSlice := strings.Fields(grades)
	updatedGrades := make([]string, 0)

	gradeSlice[gradePosition] = newGrade
	updatedGrades = append(updatedGrades, gradeSlice...)

	return strings.Join(updatedGrades, " ")
}

func UpdateFileContent(file *os.File, content []string) {
	file.Truncate(0)

	file.Seek(0, 0)

	for _, v := range content {
		file.WriteString(fmt.Sprintf("%s\n", v))
	}
}

func main() {
	file, err := os.OpenFile("student_notes.txt", os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	PrintFileContent(file)

	// Task: Update the first note of student of ID 2 from 80 to 100
	studentInfo := getStudentByID(2, file)
	student, grades := getStudentGrades(studentInfo)
	updatedGrades := updateGrade(grades, 0, "100")

	newStudentInfo := strings.Join([]string{student, updatedGrades}, " ")
	updatedContent := GetFileContentUpdated(file, 2, newStudentInfo)
	UpdateFileContent(file, updatedContent)
}
