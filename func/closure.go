package main

import "fmt"

func func1() (i int) {
	i = 10
	defer func() {
		i += 1
	}()
	return 5
}

func main() {
	closure := func1()
	fmt.Println(closure)
}
