package academy

import (
	"math"
	"testing"
)

func TestAverageGrade(t *testing.T) {
	t.Run("Empty slice", func(t *testing.T) {
		grades := []int{}
		student := Sophomore{grades: grades}
		expected := 0
		if result := student.averageGrade(); result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		}
	})

	t.Run("Single grade", func(t *testing.T) {
		grades := []int{3}
		student := Sophomore{grades: grades}
		expected := 3
		if result := student.averageGrade(); result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		}
	})

	t.Run("Multiple grades", func(t *testing.T) {
		grades := []int{2, 2, 4, 4}
		student := Sophomore{grades: grades}
		expected := 3
		if result := student.averageGrade(); result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		}
	})

	t.Run("Rounding result", func(t *testing.T) {
		grades := []int{4, 3, 2, 5}
		student := Sophomore{grades: grades}
		expected := 4
		if result := student.averageGrade(); result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		}
	})
}

func TestAttendancePercentage(t *testing.T) {
	t.Run("Empty attendance", func(t *testing.T) {
		attendance := []bool{}
		student := Sophomore{attendance: attendance}
		expected := 0.0
		epsilon := 0.001
		if result := student.attendancePercentage(); math.Abs(result-expected) > epsilon {
			t.Errorf("Expected %f but got %f", expected, result)
		}
	})

	t.Run("All present", func(t *testing.T) {
		attendance := []bool{true, true, true, true, true}
		student := Sophomore{attendance: attendance}
		expected := 1.0
		epsilon := 0.001
		if result := student.attendancePercentage(); math.Abs(result-expected) > epsilon {
			t.Errorf("Expected %f but got %f", expected, result)
		}
	})

	t.Run("All absent", func(t *testing.T) {
		attendance := []bool{false, false, false, false, false}
		student := Sophomore{attendance: attendance}
		expected := 0.0
		epsilon := 0.001
		if result := student.attendancePercentage(); math.Abs(result-expected) > epsilon {
			t.Errorf("Expected %f but got %f", expected, result)
		}
	})

	t.Run("Mixed attendance", func(t *testing.T) {
		attendance := []bool{true, true, false, true, false, false, true}
		student := Sophomore{attendance: attendance}
		expected := 0.571
		epsilon := 0.001
		if result := student.attendancePercentage(); math.Abs(result-expected) > epsilon {
			t.Errorf("Expected %f but got %f", expected, result)
		}
	})
}

func TestFinalGrade(t *testing.T) {
	testCases := []struct {
		name     string
		student  Sophomore
		expected int
	}{
		{
			name: "high attendance, high grades",
			student: Sophomore{
				name:       "John Doe",
				grades:     []int{5, 4, 5, 5, 5},
				project:    5,
				attendance: []bool{true, true, true, true, true},
			},
			expected: 5,
		},
		{
			name: "mediocre attendance, high grades",
			student: Sophomore{
				name:       "Jane Smith",
				grades:     []int{5, 4, 5, 5, 5},
				project:    5,
				attendance: []bool{true, false, true, true, false},
			},
			expected: 4,
		},
		{
			name: "low attendance, high grades",
			student: Sophomore{
				name:       "Jane Smith",
				grades:     []int{5, 4, 5, 5, 5},
				project:    4,
				attendance: []bool{true, false, false, true, false},
			},
			expected: 1,
		},
		{
			name: "mediocre attendance, low grades",
			student: Sophomore{
				name:       "Jane Smith",
				grades:     []int{2, 2, 3, 2, 2},
				project:    2,
				attendance: []bool{true, false, true, true, false},
			},
			expected: 1,
		},
		{
			name: "rounding to the nearest integer",
			student: Sophomore{
				name:       "John",
				grades:     []int{4, 4, 4},
				project:    5,
				attendance: []bool{true, true},
			},
			expected: 5,
		},
		{
			name: "high average grade, low grade from project",
			student: Sophomore{
				name:       "John Doe",
				grades:     []int{5, 4, 5, 5, 5},
				project:    2,
				attendance: []bool{true, true, true, true, true},
			},
			expected: 4,
		},
		{
			name: "high average grade, failed project",
			student: Sophomore{
				name:       "John Doe",
				grades:     []int{5, 4, 5, 5, 5},
				project:    1,
				attendance: []bool{true, true, true, true, true},
			},
			expected: 1,
		},
		{
			name: "failed semester, high grade from project",
			student: Sophomore{
				name:       "John Doe",
				grades:     []int{1, 2, 1, 1, 2},
				project:    5,
				attendance: []bool{true, true, true, true, true},
			},
			expected: 1,
		},
		{
			name: "should not decrease below minimum grade",
			student: Sophomore{
				name:       "Jane Smith",
				grades:     []int{1, 1, 1, 1, 1},
				project:    1,
				attendance: []bool{true, false, true, true, false},
			},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.student.FinalGrade()
			if result != tc.expected {
				t.Errorf("Expected %d but got %d", tc.expected, result)
			}

		})
	}
}

func TestName(t *testing.T) {
	want := "Jan Kowalski"
	got := Sophomore{name: want}.Name()
	if got != want {
		t.Errorf("Expected %s but got %s", want, got)
	}
}


func TestYear(t *testing.T) {
	got := Sophomore{}.Year()
	if got != 2 {
		t.Errorf("Expected 2 but got %d", got)
	}
}
