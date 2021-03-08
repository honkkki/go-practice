package main

import "fmt"

// 获取命令行输入
func main() {
	var a int
	var b string
	//var c float64
	//_, _ = fmt.Scanf("%d%s", &a, &b)
	//fmt.Scanf("%d", &a)
	//fmt.Scanf("%s", &b)
	//fmt.Scanf("%f", &c)
	//fmt.Println(a, b, c)

	fmt.Scan(&a, &b) // 回车换行分隔输入
	fmt.Println(a, b)

	fmt.Scanln(&a, &b) // 空格分隔输入 遇到回车退出
	fmt.Println(a, b)

}
