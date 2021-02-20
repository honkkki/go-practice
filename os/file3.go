package main

import (
	"fmt"
	"io/ioutil"
)

// ioutil

func main() {
	content, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}

	fmt.Println(string(content))
}
