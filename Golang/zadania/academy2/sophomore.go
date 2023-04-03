package academy

import "math"

type Sophomore struct {
	name       string
	grades     []int
	project    int
	attendance []bool
}

// averageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func (s Sophomore) averageGrade() int {
	if len(s.grades) == 0 {
		return 0
	}
	var sum float64
	for _, grade := range s.grades {
		sum += float64(grade)
	}
	return int(math.Round(sum / float64(len(s.grades))))
}

// attendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from 0 to 1.
func (s Sophomore) attendancePercentage() float64 {
	if len(s.attendance) == 0 {
		return 0
	}
	var sum float64
	for _, present := range s.attendance {
		if present {
			sum += 1
		}
	}
	return sum / float64(len(s.attendance))
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
func (s Sophomore) FinalGrade() int {
	attendance := s.attendancePercentage()
	avgGrade := s.averageGrade()
	grade := int(math.Round(float64(avgGrade+s.project) / 2))
	switch {
	case attendance < 0.6 || avgGrade == 1 || s.project == 1:
		return 1
	case attendance < 0.8:
		grade -= 1
	}
	return grade
}

func (s Sophomore) Name() string {
	return s.name
}

func (s Sophomore) Year() uint8 {
	return 2
}
