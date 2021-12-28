package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("连接服务器错误：", err)
		return
	}

	defer conn.Close()

	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		input = strings.TrimSpace(input)
		if strings.ToUpper(input) == "Q" {
			break
		}

		_, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println("发送数据错误: ", err)
			break
		}

		var buf [1024]byte
		n, err := conn.Read(buf[:])

		if err != nil {
			fmt.Println("从服务器获取数据失败：", err)
		}

		fmt.Println("服务器返回数据：", string(buf[:n]))
	}

}
