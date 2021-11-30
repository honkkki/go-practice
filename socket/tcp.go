package main

import (
	"fmt"
	"net"
)

type connInfo struct {
	id   int
	conn net.Conn
}

// 处理任务
func process(connInfo connInfo) {
	conn := connInfo.conn
	defer conn.Close()

	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("读取客户端数据失败: ", err)
			return
		}

		fmt.Printf("从客户端id:%d 收到数据: %v\n", connInfo.id, string(buf[:n]))
		conn.Write([]byte("服务器已收到数据"))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:6333")
	if err != nil {
		fmt.Printf("start tcp failed: %v \n", err)
		return
	}

	var id int

	for {
		id++
		conn, err := listen.Accept()
		connInfo := connInfo{
			id:   id,
			conn: conn,
		}

		if err != nil {
			fmt.Printf("conn failed: %v \n", err)
			continue
		}

		go process(connInfo)
	}
}
