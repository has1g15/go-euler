package GradingStudents

func gradingStudent(grade int) int {
	if grade > 37 {
		plusOne := grade + 1
		if plusOne % 5 == 0 {
			grade = plusOne
		}
		plusTwo := grade + 2
		if plusTwo % 5 == 0 {
			grade = plusTwo
		}
	}
	return grade
}
