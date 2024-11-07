package file_handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PrintFileContent(file *os.File) {
	fileContent := GetFileContent(file)

	for _, v := range fileContent {
		fmt.Println(v)
	}
}

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

func UpdateFileContent(file *os.File, content []string) {
	file.Truncate(0)

	file.Seek(0, 0)

	for _, v := range content {
		file.WriteString(fmt.Sprintf("%s\n", v))
	}
}

func AppendToFile(file *os.File, content string) {
	file.WriteString(fmt.Sprintf("\n%s", content))
}
