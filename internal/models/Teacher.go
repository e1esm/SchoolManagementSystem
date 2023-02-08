package models

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
	name           string
	CurrentSubject Subject
	id             int
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
