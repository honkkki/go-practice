package main

import (
	"fmt"
	"reflect"
)

type Singer struct {
	Name string
	Age  int
}

func main() {
	var s Singer
	v := reflect.ValueOf(&s)
	v.Elem().Field(0).SetString("jisoo")
	v.Elem().FieldByName("Age").SetInt(20)
	fmt.Println(s)

}
