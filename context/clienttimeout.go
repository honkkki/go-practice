// 模拟http服务超时控制
package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

type ResPack struct {
	res *http.Response
	err error
}

func req(ctx context.Context) {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr, Timeout: time.Second * 5}
	defer wg.Done()
	c := make(chan ResPack, 1)
	req, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	// 开启一个goroutine去请求http server 通过channel有没有值判断请求是否完成
	go func(c chan<- ResPack) {
		resp, err := client.Do(req)
		pack := ResPack{res: resp, err: err}
		c <- pack
	}(c)

	select {
	case <-ctx.Done():
		// 超时请求
		//tr.CancelRequest(req)
		fmt.Println("http timeout")
	case data := <-c:
		if data.res == nil {
			fmt.Println("null resp")
			return
		}
		// 正常请求
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(data.res.Body)

		if data.err != nil {
			fmt.Println(data.err)
			return
		}
		r, _ := ioutil.ReadAll(data.res.Body)
		fmt.Printf("server response: %s", r)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	wg.Add(1)
	go req(ctx)
	wg.Wait()
	fmt.Println("finish")
}
