package main

import (
	"bufio"
	"fmt"
	"net"
)

// 处理任务
func process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [200]byte
		n, err := reader.Read(buf[:])

		if err != nil {
			fmt.Println("读取客户端数据失败: ", err)
			return
		}

		fmt.Println("从客户端收到的数据：", string(buf[:n]))
		conn.Write([]byte("服务器已收到数据"))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:6333")
	if err != nil {
		fmt.Printf("start tcp failed: %v \n", err)
		return
	}

	for {
		conn, err := listen.Accept()

		if err != nil {
			fmt.Printf("conn failed: %v \n", err)
			continue
		}

		go process(conn)
	}
}
