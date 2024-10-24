package students

type Student struct {
	ID     int
	Name   string
	Grades []int
}

func (s *Student) AddGrade(grade int) {
	s.Grades = append(s.Grades, grade)
}

func (s *Student) GetAverage() int {
	studentGrades := s.Grades
	if len(studentGrades) == 0 {
		return 0
	}

	var sumUpGrades int
	for _, grade := range studentGrades {
		sumUpGrades = sumUpGrades + grade
	}
	average := sumUpGrades / len(studentGrades)
	return average
}
