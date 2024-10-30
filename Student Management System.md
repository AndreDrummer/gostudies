# Go Studies

### `1` GO Practice Challenge: Student Management System

You are tasked with creating a simple Student Management System using Go. The program will allow users to add students, remove students, and display a list of students, including their grades, average scores, and whether they passed or failed based on a minimum average.

### **Requirements:**

1. **Define a `Student` struct**:
   - Each `Student` should have the following properties:
     - `Name` (string)
     - `Grades` (slice of integers)

2. **Create a map to store the students**:
   - Use a map with the `Name` of the student as the key and the `Student` struct as the value.

3. **Functions**:
   - Create a function `addStudent(name string)` that adds a new student with no grades to the map.
   - Create a function `addGrade(name string, grade int)` that adds a grade to the student’s list of grades.
   - Create a function `removeStudent(name string)` that removes a student by name from the map.
   - Create a function `calculateAverage(name string)` that calculates and returns the average of the student’s grades.
   - Create a function `checkPassOrFail(name string)` that returns if the student passed or failed. The passing average is 60.

4. **Main Menu**:
   - In the `main` function, create a menu with the following options:
     - `1` - Add a new student.
     - `2` - Add a grade to a student.
     - `3` - Remove a student.
     - `4` - Calculate and display the average score of a student.
     - `5` - Display whether a student passed or failed.
     - `6` - Display all students and their grades.
     - `0` - Exit the program.

5. **Bonus**:
   - Allow users to input multiple grades at once when adding grades.
   - Add input validation for names and grades (e.g., a grade should be between 0 and 100).

---

### **Example interaction**:

```
Welcome to the Student Management System!
1 - Add a new student
2 - Add a grade to a student
3 - Remove a student
4 - Calculate average score of a student
5 - Check if a student passed or failed
6 - Display all students and their grades
0 - Exit

Enter your choice: 1
Enter student's name: Alice
Student Alice added!

Enter your choice: 2
Enter student's name: Alice
Enter grade: 85
Grade 85 added to Alice!

Enter your choice: 4
Enter student's name: Alice
Alice's average score: 85.0

Enter your choice: 5
Enter student's name: Alice
Alice passed!

Enter your choice: 6
List of students:
- Alice: [85]
```