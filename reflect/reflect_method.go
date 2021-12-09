package main

import (
	"fmt"
	"reflect"
)

type St struct {
	Name string
}

func (s *St) Test() {
	fmt.Println("i am reflect test function")
}

func (s *St) Add(x int) int {
	x++
	return x
}

func main() {
	s := new(St) // 注意方法定义中的接收者为指针类型
	v := reflect.ValueOf(s)
	t := v.Type()

	m := t.Method(0)
	fmt.Println(m)
	fmt.Println(m.Name)
	fmt.Println(m.Type)

	fmt.Println("---------------------")
	fmt.Println(v.Method(0))
	v.MethodByName("Test").Call([]reflect.Value{})

	val := v.MethodByName("Add").Call([]reflect.Value{
		reflect.ValueOf(1),
	})

	fmt.Println(len(val))
	fmt.Println(val[0].Interface())

}
