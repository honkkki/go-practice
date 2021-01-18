package main

import "fmt"

func main()  {
	i := 0
	defer fmt.Println(i)		// i=0 已经赋值
	i = 1000
	fmt.Println(i)

}
