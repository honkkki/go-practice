package main

import "fmt"

type Addable interface {
	type int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64, complex64, complex128,
		string
}

func print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func add[T Addable] (a, b T) T {
	return a + b
}

func main() {
	s1 := []string{"golang", "hello"}
	s2 := []int{1, 2, 3}
	print(s1)
	print(s2)

	fmt.Println(add(1, 2))
	fmt.Println(add(1.1, 2.2))
	fmt.Println(add("hello ", "golang"))
}
