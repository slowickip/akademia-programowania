package academy

type Repository interface {
	List(year uint8) (names []string, err error)
	Get(name string) (Student, error)
	Save(name string, year uint8) error
	Graduate(name string) error
}
