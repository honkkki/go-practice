package main

import "fmt"

type Singer struct {
	name string
	age  int
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
