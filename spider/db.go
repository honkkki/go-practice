package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var i interface{}
	i = 1
	fmt.Println(reflect.TypeOf(i).Kind())

}


