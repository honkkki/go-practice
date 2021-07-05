// 接口是可比较的 类型与值的比较
package main

import "fmt"

type Monitor interface {
	show(content string) string
}

type Mouse interface {
	click()
}

type Acer struct {
	Name string
}

func (a *Acer) show(content string) string {
	return fmt.Sprintf("%s : %s", a.Name, content)
}


func main()  {
	var a interface{} = 1
	var b interface{} = 1
	fmt.Println(a==b)
	fmt.Println("----------------------------------------")

	var m Monitor
	var m2 Monitor
	fmt.Println(m==m2)		// both nil true
	m = &Acer{Name: "acer"}
	ac := m.(*Acer)
	fmt.Println(m==ac)		// true
	fmt.Printf("%T\n", m)
	fmt.Printf("%T\n", ac)
	m2 = &Acer{Name: "acer"}
	fmt.Println(m==m2)		// false 指针地址不同 如果存的是结构体是相等的
	fmt.Println("----------------------------------------")

	fmt.Println(m.show("hi"))

	var arr interface{}
	arr = []int{1,2,3}
	fmt.Println(arr)
	//fmt.Println(arr[1])		// 不能这样取值

	var num interface{}
	num = 1
	fmt.Println(1==num)			// true 比较是相等的 但无法当做int类型使用 需要断言


}
