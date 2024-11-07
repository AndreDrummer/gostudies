package main

import system_initializer "github.com/AndreDrummer/gostudies/StudentManagementSystem/System/initializer"

func main() {
	system_initializer.Initialize()
	// file, err := os.OpenFile("student_notes.txt", os.O_APPEND|os.O_WRONLY, 0644)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// file.WriteString("\nTeste")

}
