package models

import (
	"github.com/jackc/pgx"
	"log"
)

type Subject string

const (
	Maths           Subject = "Maths"
	ForeignLanguage Subject = "English"
	Informatics     Subject = "Informatics"
	NativeLanguage  Subject = "Native language"
	Drawing         Subject = "Drawing"
	Music           Subject = "Music"
	Literature      Subject = "Literature"
	LawScience      Subject = "Law Science"
)

type HeadTeacher interface {
	SetMark(int, string)
	AddStudent(Learner)
	MarksToPdfOf(string)
	AllMarksToPDF() error
}

type Teacher struct {
	Name           string
	CurrentSubject Subject
	ID             int
}

func (t Teacher) SetMark(grade int, studentName string) {

}
func (t Teacher) AddStudent(student Student) {

}

func (t Teacher) MarksToPdF(subject Subject) {

}

func (t Teacher) AllMarksToPDF() error {

	return nil
}

func TeacherGenerator(guest Guest, subject Subject) Teacher {
	return Teacher{Name: guest.Username, CurrentSubject: subject}
}

func (s Student) leave(db *pgx.Conn) {
	sqlDeleteQuery := "UPDATE users set is_left = true where name = $1"
	_, err := db.Exec(sqlDeleteQuery, s.name)
	if err != nil {
		log.Fatal("Couldn't leave the school as a teacher")
	}
}
