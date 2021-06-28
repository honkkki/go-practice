package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConnEcho(c net.Conn) {
	input := bufio.NewScanner(c)

	// 阻塞监听客户端输入直到客户端断开
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}

	fmt.Println("close server")
	c.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err) // e.g., connection aborted
			continue
		}

		handleConnEcho(conn)
	}
}
