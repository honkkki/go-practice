package main

import (
	"fmt"
	"runtime"
)

func main()  {
	// 查看系统cpu数目 go默认使用全部cpu数
	fmt.Println(runtime.NumCPU())


}
