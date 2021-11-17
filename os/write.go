package main

import (
	"bufio"
	"fmt"
	"os"
)

// bufio写入文件
func main() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}
	defer file.Close()

	//bufio.NewWriterSize(file, 256)
	writer := bufio.NewWriter(file)
	writer.WriteString("hello golang\n")
	writer.Flush()
}
