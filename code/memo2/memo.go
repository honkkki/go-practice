package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

// map的value 存的是url请求的结果与是否有缓存的flag
type entry struct {
	res   result
	ready chan struct{}
}

// 请求对象 存url、请求的响应结果channel
type request struct {
	key      string
	response chan<- result
}

type Memo struct {
	requests chan request
}

type result struct {
	value interface{}
	err   error
}

func NewMemo(f Func) *Memo {
	memo := &Memo{requests: make(chan request, 10)}
	go memo.Server(f)
	return memo
}

// Get 获取响应的结果
// 每个请求封装为一个request对象放进管道里
// 等到下面的deliver方法获取到结果后返回结果
func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{
		key:      key,
		response: response,
	}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

// Server monitor goroutine
// 处理所有的url请求 消费request管道
func (memo *Memo) Server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		// 把响应结果写进对应的管道里
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	fmt.Println("http get")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	m := NewMemo(httpGetBody)
	defer m.Close()
	incomingURLs := []string{
		"https://golang.org",
		"https://www.baidu.com",
		"https://gobyexample.com",
		"https://books.studygolang.com",
		"https://golang.org",
		"https://gobyexample.com",
		"https://golang.org",
		"https://golang.org",
	}

	var wg sync.WaitGroup

	for _, url := range incomingURLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Printf("get data fail: %v\n", err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}

	wg.Wait()
}
