package main

import (
	"errors"
	"fmt"
)

func main()  {
	// error
	err := errors.New("foo")
	fmt.Printf("%T %v \n", err, err)
	fmt.Println(err.Error())
}
