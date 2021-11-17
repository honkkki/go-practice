package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.OpenFile("./test2.txt", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}
	defer file.Close()

	//file.WriteString("hello gogogo\n")

	// ioutil读写文件
	content, _ := ioutil.ReadFile("./test.txt")
	ioutil.WriteFile("./test3.txt", content, 0777)

}
