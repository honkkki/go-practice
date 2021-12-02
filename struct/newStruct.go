package main

import (
	"fmt"
)

// 接口嵌套

type Writer interface {
}

type Reader interface {
	Writer
}

type Singer struct {
	name string
	age  int
	Writer
}

func newSinger(name string, age int) *Singer {
	return &Singer{
		name: name,
		age:  age,
	}
}

func main() {
	singer := newSinger("jisoo", 20)
	fmt.Printf("%#v \n", singer)
	fmt.Println(singer.name)

}
