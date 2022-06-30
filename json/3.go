package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

type Test struct {
	Name string `json:"name"`
}

func main() {
	data := Test{}
	str := `{"name":null}`
	jsoniter.UnmarshalFromString(str, &data)
	fmt.Println(data)
}
