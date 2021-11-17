package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// bufio读取文件
func main() {
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		fmt.Println(line)
		if err == io.EOF {
			break
		}
	}

}
