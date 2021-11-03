package main

import "fmt"

func main() {
	defer func() {
		if i := recover(); i != nil {
			fmt.Println("recover:", i)
		}
	}()

	panic("something panic.")
}
