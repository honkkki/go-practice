package main

import "fmt"

func main() {
	m := make(map[string]interface{})
	m["name"] = "tom"
	m["age"] = 1
	fmt.Println(len(m))

	for k, v := range m {
		fmt.Println(v)
		delete(m, k)
	}

	fmt.Println(len(m))

}
