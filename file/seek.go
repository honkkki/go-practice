package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("1.go")
	if err != nil {
		log.Fatal(err)
	}

	// 通过seek设置光标位置读写文件
	buf := make([]byte, 1, 1)
	f.Read(buf)
	fmt.Println(string(buf))
	f.Seek(1, io.SeekStart)
	f.Read(buf)
	fmt.Println(string(buf))

	f.Seek(0, io.SeekCurrent)
	f.Read(buf)
	fmt.Println(string(buf))

}
