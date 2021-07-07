package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
)

type client chan string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

// broadcaster 其他goroutine通过操作channel来改变map 且只有一个goroutine执行该函数 没有并发竞争问题
func broadcaster(ctx context.Context, wg *sync.WaitGroup) {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli) // 关闭下面的ch
		case <-ctx.Done():
			fmt.Println("broadcaster quit")
			wg.Done()
			return
		}
	}
}

func handleConn(conn net.Conn) {
	cli := make(chan string) // outgoing client messages
	go clientWriter(conn, cli)

	who := conn.RemoteAddr().String()
	cli <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- cli // 在broadcaster方法中关闭
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:9999")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	// 监听命令退出程序
	// 退出程序时退出broadcaster goroutine
	go func() {
		var wg sync.WaitGroup
		ctx, cancel := context.WithCancel(context.Background())
		go broadcaster(ctx, &wg)

		for {
			select {
			case <-sig:
				defer listener.Close()
				wg.Add(1)
				fmt.Println("server quit")
				cancel()
				wg.Wait()
				os.Exit(1)
			}
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listen error: %v", err)
			continue
		}
		go handleConn(conn)
	}
}
