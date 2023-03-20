package main

import (
	"github.com/google/uuid"
	"github.com/grupawp/appdispatcher"
	"log"
)

type Student struct {
	FirstName     string
	LastName      string
	applicationID uuid.UUID
}

func (s Student) ApplicationID() string {
	str := s.applicationID.String()
	return str
}

func (s Student) FullName() string {
	fullName := s.FirstName + " " + s.LastName
	return fullName
}

func main() {
	student := Student{
		FirstName:     "Piotr",
		LastName:      "Słowicki",
		applicationID: uuid.New(),
	}
	submit, err := appdispatcher.Submit(student)
	log.Default().Println(submit)
	log.Default().Println(err)
}
