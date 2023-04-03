package academy

import "github.com/pkg/errors"

var ErrInvalidGrade = errors.New("invalid grade")
var ErrStudentNotFound = errors.New("student not found")

func GradeYear(r Repository, year uint8) error {
	students, err := r.List(year)
	if err != nil {
		return err
	}

	for _, s := range students {
		err = GradeStudent(r, s)
		if err != nil {
			return err
		}
	}

	return nil
}

func GradeStudent(r Repository, name string) error {
	s, err := r.Get(name)
	switch {
	case errors.Is(ErrStudentNotFound, err):
		return nil
	case err != nil:
		return err
	}

	grade := s.FinalGrade()
	if grade < 1 || grade > 5 {
		return ErrInvalidGrade
	}

	switch {
	case grade == 1:
		err = r.Save(s.Name(), s.Year())
	case s.Year() == 3:
		err = r.Graduate(s.Name())
	default:
		err = r.Save(s.Name(), s.Year()+1)
	}

	return err
}
