package academy

import (
	"math"
)

type Student struct {
	Name      string
	Grades    []int
	Project   int
	Attendace []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	//panic("not implemented")

	if len(grades) != 0 {
		sum := 0
		ctr := 0

		for range grades {
			sum += grades[ctr]
			ctr++
		}

		return int(math.Round(float64(sum) / float64(ctr)))
	} else {
		return 0
	}
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from  0 to 1,
// with 2 digits of precision.
func AttendancePercentage(attendance []bool) float64 {
	//panic("not implemented")

	if len(attendance) != 0 {
		var attendanceSum int

		for counter := 0; counter < len(attendance); counter++ {

			if attendance[counter] {
				attendanceSum++
			}
		}

		return math.Round(float64(attendanceSum)/float64(len(attendance))*1000) / 1000
	} else {
		return 0
	}
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.

// testy wymagają 3 miejsc po przecinku
func FinalGrade(s Student) int {
	//panic("not implemented")

	avgGrades := AverageGrade(s.Grades)
	attendance := AttendancePercentage(s.Attendace)

	switch {
	case s.Project == 1 || avgGrades == 1:
		return 1
	case attendance < 0.6:
		return 1
	case attendance < 0.8:
		return int(math.Round(float64(avgGrades+s.Project)/2)) - 1
	default:
		return int(math.Round(float64(avgGrades+s.Project) / 2))
	}
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	//panic("not implemented")

	grades := make(map[string]uint8)

	for i := 0; i < len(students); i++ {
		grades[students[i].Name] = uint8(FinalGrade(students[0]))
	}

	return grades
}
