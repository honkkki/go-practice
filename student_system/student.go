package main

type Student struct {
	ID   int
	Name string
}

func NewStudent(id int, name string) *Student {
	return &Student{
		ID:   id,
		Name: name,
	}
}
