package GradingStudents

import "testing"

func TestGradingStudents(t *testing.T) {
	var tests = []struct {
		input int
		expected int
	} {
		{75, 75},
		{67, 67},
		{38, 40},
		{33, 33},
	}

	for _, test := range tests {
		if result := gradingStudent(test.input); result != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received {}", test.input, test.expected, result)
		}
	}
}
