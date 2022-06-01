package main

import "fmt"

type Base struct {
	Name string
	Num  int64
}

func main() {
	source := []Base{
		{
			Name: "tom",
			Num:  1,
		},
		{
			Name: "tim",
			Num:  1,
		},
		{
			Name: "tom",
			Num:  2,
		},
	}

	var source2 = []Base{
		{
			Name: "tom",
			Num:  1,
		},
	}
	var res []Base
	res = append(source2, source...)
	fmt.Println(res)
}
