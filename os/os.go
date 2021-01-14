package main

import (
	"fmt"
	"os"
	"strings"
)

// 获取终端输入
func main()  {
	var s string
	fmt.Println(os.Args)

	for i:=1;i<len(os.Args);i++ {
		s = s + os.Args[i] + " "
	}

	s2 := strings.Join(os.Args[1:], " ")


	fmt.Println(s)
	fmt.Println(s2)
}
