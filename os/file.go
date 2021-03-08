package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}

	defer file.Close()

	var buf [64]byte

	n, err := file.Read(buf[:])
	//if err == io.EOF {
	//
	//}

	fmt.Println(n)
	fmt.Println(string(buf[:n]))

}
