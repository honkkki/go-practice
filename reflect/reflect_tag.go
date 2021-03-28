package main

import (
	"fmt"
	"reflect"
)

type Game struct {
	Name string `json:"name"`
}

func main()  {
	g := Game{Name: "rainbow"}
	t := reflect.TypeOf(&g)

	t1 := t.Elem().Field(0).Tag.Get("json")
	fmt.Println(t1)
}
