package models

type Learner interface {
	GetAllMarks() []int
	GetMarksFor(string) []int
}

type Student struct {
	name string
	age  int
	id   int
}

func (s Student) GetAllMarks() []int {

	return nil
}

func (s Student) getMarksFor(subject Subject) []int {
	return nil
}

func StudentGenerator(age int, name string) *Student {
	return &Student{age: age, name: name}
}
