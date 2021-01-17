package main

import (
	"fmt"
	"time"
)

func main()  {
	now := time.Now()
	fmt.Println(now)
	// 时间戳
	fmt.Println(now.Unix())

	year := now.Year()
	month := now.Month()
	day := now.Day()

	// 格式化输出
	fmt.Printf("%02d-%02d-%02d \n", year, month, day)
	nowTime := time.Unix(now.Unix(), 0)
	fmt.Println(nowTime.Year(), nowTime.Month())

}
