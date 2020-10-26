package main

import (
	"encoding/json"
	"fmt"
)

// json序列化示例
type Student struct {
	ID   int
	Name string
}

type Class struct {
	Name     string    `json:"name"`
	Students []Student `json:"students"`
}

func newStudent(id int, name string) *Student {
	return &Student{
		ID:   id,
		Name: name,
	}
}

func main() {
	c1 := Class{
		Name:     "三年二班",
		Students: make([]Student, 0, 10),
	}

	for i := 0; i < 10; i++ {
		stu := newStudent(i, fmt.Sprintf("student%02d", i))
		c1.Students = append(c1.Students, *stu)
	}

	fmt.Printf("%#v \n", c1)

	data, err := json.Marshal(c1)
	if err != nil {
		panic("json encode error")
	}

	fmt.Printf("json序列化后: %v \n", string(data))

	var c2 Class
	err = json.Unmarshal(data, &c2)

	if err != nil {
		panic("json decode error")
	}

	fmt.Printf("json反序列化: %#v \n", c2)
}
