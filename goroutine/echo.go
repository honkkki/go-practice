// gopl 8.8
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
}

func handleConnEcho(c net.Conn) {
	input := bufio.NewScanner(c)
	message := make(chan string)
	var wg sync.WaitGroup

	go func() {
		for {
			ctx, cancel := context.WithCancel(context.Background())
			fmt.Println("server down after 10 seconds!")
			go func(ctx context.Context) {
				tick := time.NewTicker(1 * time.Second)
				defer tick.Stop()
				for i := 10; i > 0; i-- {
					select {
					case <-ctx.Done():
						// 重置倒计时
						fmt.Println("reset deadline")
						return
					default:
						fmt.Println(i)
						<-tick.C
					}
				}
			}(ctx)

			timer := time.NewTimer(10 * time.Second)
			select {
			case <-timer.C:
				fmt.Println("time to close connection")
				c.Close()
				cancel()
				return
			case msg := <-message:
				wg.Add(1)
				go echo(c, msg, 1*time.Second, &wg)
				wg.Wait()
				timer.Stop()
				cancel()
			}
		}
	}()

	// 阻塞监听客户端输入直到客户端断开
	for input.Scan() {
		//go echo(c, input.Text(), 1*time.Second)
		message <- input.Text()
	}

	fmt.Println("client close connection")
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
			log.Println(err)
			continue
		}

		handleConnEcho(conn)
	}
}
