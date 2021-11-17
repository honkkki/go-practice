package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var f float64 = 3.14
	str := strconv.FormatFloat(f, 'f', 2, 64)
	fmt.Println(str)
	fmt.Println(reflect.TypeOf(str))


}
