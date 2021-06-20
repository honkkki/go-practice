package main

import "fmt"

func main()  {
	var i interface{}
	i = make(map[string]interface{})
	fmt.Printf("%T\n", i)

	m := i.(map[string]interface{})
	m["name"] = "jisoo"
	fmt.Println(m)
	fmt.Printf("%T\n", m)

}
