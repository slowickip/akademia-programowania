package academy

import (
	"math"
	"reflect"
	"testing"
)

func TestAverageGrade(t *testing.T) {
	t.Run("Empty slice", func(t *testing.T) {
		grades := []int{}
		expected := 0
		if result := AverageGrade(grades); result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		}
	})

	t.Run("Single grade", func(t *testing.T) {
		grades := []int{3}
		expected := 3
		if result := AverageGrade(grades); result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		}
	})

	t.Run("Multiple grades", func(t *testing.T) {
		grades := []int{2, 2, 4, 4}
		expected := 3
		if result := AverageGrade(grades); result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		}
	})

	t.Run("Rounding result", func(t *testing.T) {
		grades := []int{4, 3, 2, 5}
		expected := 4
		if result := AverageGrade(grades); result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		}
	})
}

func TestAttendancePercentage(t *testing.T) {
	t.Run("Empty attendance", func(t *testing.T) {
		attendance := []bool{}
		expected := 0.0
		epsilon := 0.001
		if result := AttendancePercentage(attendance); math.Abs(result-expected) > epsilon {
			t.Errorf("Expected %f but got %f", expected, result)
		}
	})

	t.Run("All present", func(t *testing.T) {
		attendance := []bool{true, true, true, true, true}
		expected := 1.0
		epsilon := 0.001
		if result := AttendancePercentage(attendance); math.Abs(result-expected) > epsilon {
			t.Errorf("Expected %f but got %f", expected, result)
		}
	})

	t.Run("All absent", func(t *testing.T) {
		attendance := []bool{false, false, false, false, false}
		expected := 0.0
		epsilon := 0.001
		if result := AttendancePercentage(attendance); math.Abs(result-expected) > epsilon {
			t.Errorf("Expected %f but got %f", expected, result)
		}
	})

	t.Run("Mixed attendance", func(t *testing.T) {
		attendance := []bool{true, true, false, true, false, false, true}
		expected := 0.571
		epsilon := 0.001
		if result := AttendancePercentage(attendance); math.Abs(result-expected) > epsilon {
			t.Errorf("Expected %f but got %f", expected, result)
		}
	})
}

func TestFinalGrade(t *testing.T) {
	testCases := []struct {
		name     string
		student  Student
		expected int
	}{
		{
			name: "high attendance, high grades",
			student: Student{
				Name:       "John Doe",
				Grades:     []int{5, 4, 5, 5, 5},
				Project:    5,
				Attendance: []bool{true, true, true, true, true},
			},
			expected: 5,
		},
		{
			name: "mediocre attendance, high grades",
			student: Student{
				Name:       "Jane Smith",
				Grades:     []int{5, 4, 5, 5, 5},
				Project:    5,
				Attendance: []bool{true, false, true, true, false},
			},
			expected: 4,
		},
		{
			name: "low attendance, high grades",
			student: Student{
				Name:       "Jane Smith",
				Grades:     []int{5, 4, 5, 5, 5},
				Project:    4,
				Attendance: []bool{true, false, false, true, false},
			},
			expected: 1,
		},
		{
			name: "mediocre attendance, low grades",
			student: Student{
				Name:       "Jane Smith",
				Grades:     []int{2, 2, 3, 2, 2},
				Project:    2,
				Attendance: []bool{true, false, true, true, false},
			},
			expected: 1,
		},
		{
			name: "rounding to the nearest integer",
			student: Student{
				Name:       "John",
				Grades:     []int{4, 4, 4},
				Project:    5,
				Attendance: []bool{true, true},
			},
			expected: 5,
		},
		{
			name: "high average grade, low grade from project",
			student: Student{
				Name:       "John Doe",
				Grades:     []int{5, 4, 5, 5, 5},
				Project:    2,
				Attendance: []bool{true, true, true, true, true},
			},
			expected: 4,
		},
		{
			name: "high average grade, failed project",
			student: Student{
				Name:       "John Doe",
				Grades:     []int{5, 4, 5, 5, 5},
				Project:    1,
				Attendance: []bool{true, true, true, true, true},
			},
			expected: 1,
		},
		{
			name: "failed semester, high grade from project",
			student: Student{
				Name:       "John Doe",
				Grades:     []int{1, 2, 1, 1, 2},
				Project:    5,
				Attendance: []bool{true, true, true, true, true},
			},
			expected: 1,
		},
		{
			name: "should not decrease below minimum grade",
			student: Student{
				Name:       "Jane Smith",
				Grades:     []int{1, 1, 1, 1, 1},
				Project:    1,
				Attendance: []bool{true, false, true, true, false},
			},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FinalGrade(tc.student)
			if result != tc.expected {
				t.Errorf("Expected %d but got %d", tc.expected, result)
			}

		})
	}
}

func TestGradeStudents(t *testing.T) {
	t.Run("Empty students slice", func(t *testing.T) {
		students := []Student{}
		expected := map[string]uint8{}
		if result := GradeStudents(students); !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v but got %v", expected, result)
		}
	})

	t.Run("Multiple students", func(t *testing.T) {
		students := []Student{
			{
				Name:       "John",
				Grades:     []int{4, 5, 4, 3, 4},
				Project:    5,
				Attendance: []bool{true, true, false, true, true, true},
			},
			{
				Name:       "Jane",
				Grades:     []int{5, 5, 5, 5, 5},
				Project:    5,
				Attendance: []bool{true, true, true, true, true, true},
			},
		}
		expected := map[string]uint8{
			"John": 5,
			"Jane": 5,
		}
		if result := GradeStudents(students); !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v but got %v", expected, result)
		}
	})
}
