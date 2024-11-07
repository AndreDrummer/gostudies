package file_handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	utils "github.com/AndreDrummer/gostudies/StudentManagementSystem/Utils"
)

func OpenFileWithPerm(filename string, flag int) *os.File {
	file, err := os.OpenFile(filename, flag, 0644)

	if err != nil {
		fmt.Printf("ERROR %v opening file %v", err, filename)
		return nil
	}

	return file
}

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
		if line != "" {
			content = append(content, line)
		}
	}
	return content
}

func GetFileEntryByPrefix(prefix any, file *os.File) string {
	fileContent := GetFileContent(file)

	for _, v := range fileContent {
		if strings.HasPrefix(v, fmt.Sprintf("%v.", prefix)) {
			return v
		}
	}

	return ""
}

func UpdateFileEntry(file *os.File, entryPrefix any, updatedEntry string) {
	fileContent := GetFileContent(file)
	var newContent []string

	for _, v := range fileContent {
		if strings.HasPrefix(v, fmt.Sprintf("%d.", entryPrefix)) {
			newContent = append(newContent, updatedEntry)
		} else if v == "" {
			continue
		} else {
			newContent = append(newContent, v)
		}
	}

	OverrideFileContent(file, newContent)
}

func RemoveFileEntry(file *os.File, entryPrefix any) {
	fileContent := GetFileContent(file)
	var newContent []string

	for _, v := range fileContent {
		if strings.HasPrefix(v, fmt.Sprintf("%d.", entryPrefix)) || v == "" {
			continue
		} else {
			newContent = append(newContent, v)
		}
	}

	OverrideFileContent(file, newContent)
}

func OverrideFileContent(file *os.File, content []string) {
	file.Truncate(0)

	file.Seek(0, 0)

	utils.SortSliceStringByID(content, ".")

	for _, v := range content {
		file.WriteString(fmt.Sprintf("%s\n", v))
	}
}

func AppendToFile(file *os.File, content string) {
	file.WriteString(fmt.Sprintf("\n%s", content))
}

func ClearFileContent(file *os.File) {
	file.Truncate(0)
}
