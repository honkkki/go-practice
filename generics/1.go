package main

import "fmt"

type Addable interface {
	int | int8 | int16 | int32 | int64
}

func GenPrint[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func add[T Addable](a, b T) T {
	return a + b
}

// 闭包
func anony[T Addable](a, b T) func() T {
	return func() T {
		r := a
		a = a + b
		return r
	}
}

func main() {
	s1 := []string{"golang", "hello"}
	s2 := []int64{1, 2, 3}
	GenPrint(s1)
	GenPrint(s2)

	fmt.Println(add(1, 2))

	f := anony(0, 1)
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
