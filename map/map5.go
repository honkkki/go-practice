package main

import "fmt"

func main() {
	// map的value还是map的例子
	studentMap := make(map[int]map[string]interface{}, 10)

	value := make(map[string]interface{}, 10)
	value["name"] = "jisoo"
	value["age"] = 25
	studentMap[1] = value

	fmt.Println(studentMap)

}
