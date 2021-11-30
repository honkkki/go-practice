package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	msg := "GET / HTTP/1.1\r\n"
	msg += "Host:www.baidu.com\r\n"
	msg += "Connection:keep-alive\r\n"
	msg += "\r\n\r\n"

	_, err = conn.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println(string(buf[:n]))
	}
}
