package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
}

func main()  {
	var s Student
	s.Name = "jisoo"
	s.Age = 20

	v := reflect.ValueOf(s)
	t := v.Type()
	fmt.Println(v.Kind())		// struct


	fmt.Println(t.NumField())
	fmt.Println(v.NumField())

	// 根据下标获取结构体字段
	fmt.Println(t.Field(0).Name)		// 结构体属性名
	field1 := v.Field(0)		// 属性值信息
	fmt.Println(field1.Kind(), field1.Interface())	// 结构体的值信息
	name := field1.Interface()			// 返回字段的值
	fmt.Printf("name type is: %T\n", name)

}
