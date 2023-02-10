package models

import (
	"github.com/jackc/pgx"
	"log"
)

type Learner interface {
	GetAllMarks() []int
	GetMarksFor(string) []int
}

type Student struct {
	Name string
	ID   int
}

func (s Student) GetAllMarks() []int {

	return nil
}

func (s Student) getMarksFor(subject Subject) []int {
	return nil
}

func StudentGenerator(ID int, name string) *Student {
	return &Student{Name: name, ID: ID}
}

func (s Student) Leave(db *pgx.Conn) {
	sqlDeleteQuery := "UPDATE users set is_left = true where name = $1"
	_, err := db.Exec(sqlDeleteQuery, s.Name)
	if err != nil {
		log.Fatal("Couldn't leave the school as a student")
	}
}
