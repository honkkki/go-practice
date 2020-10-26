package main

import "fmt"

// 简单的学生管理系统
func main() {

	for {
		var input int
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("error")
		}

		fmt.Println(input)

		if input == 0 {
			break
		}
	}
}
