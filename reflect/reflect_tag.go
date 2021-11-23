package main

import (
	"fmt"
	"reflect"
)

type Game struct {
	Name string `json:"name"`
}

type P struct {
	name string
}

func main() {
	g := Game{Name: "rainbow"}
	t := reflect.TypeOf(&g)

	t1 := t.Elem().Field(0).Name
	t2 := t.Elem().Field(0).Tag.Get("json")
	fmt.Println(t1)
	fmt.Println(t2)

	// 反射获取结构体未导出属性
	p := P{name: "tom"}
	tt := reflect.TypeOf(p)
	fmt.Println(tt.Field(0).Name)
}
