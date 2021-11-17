package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	inputReader *bufio.Reader
)

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n') // 读取到\n结束
	if err == nil {
		fmt.Println("input:", input)
	}

}
