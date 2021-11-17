// 令牌桶限流
package main

import (
	"fmt"
	"github.com/juju/ratelimit"
	"time"
)

func main() {
	// 每秒接受100个请求
	b := ratelimit.NewBucketWithQuantum(time.Second, 100, 100)
	pre := time.Now()
	for i := 0; i < 1000; i++ {
		before := b.Available()
		tokenGet := b.TakeAvailable(1)
		if tokenGet != 0 {
			fmt.Println("获取到令牌 index=", i+1, "前后数量-> 前：", before, ", 后: ", b.Available(), ", tokenGet=", tokenGet)
		} else {
			fmt.Println("未获取到令牌，拒绝请求", i+1)
		}
		time.Sleep(1 * time.Millisecond)
	}

	after := time.Now().Sub(pre)
	fmt.Println("used time: ", after)
}
