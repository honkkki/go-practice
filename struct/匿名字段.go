package main

import (
	"fmt"
	"go-practice/struct/user"
)

type Addr struct {
	name        string
	Province    string
	City        string
	CreatedTime int
}

type Email struct {
	Url         string
	CreatedTime int
}

type Emp struct {
	name string
	Addr
	Email
}

//func (e Emp) Get()  {
//	fmt.Println("i am emp get")
//}

func (e Email) Get() {
	fmt.Println("i am email get")
}

func main() {
	var u user.User
	u.Name = "jisoo"

	fmt.Println(u)

	emp := Emp{
		name: "jisoo",
		Addr: Addr{
			name:        "addr name",
			Province:    "gd",
			City:        "sz",
			CreatedTime: 1,
		},
		Email: Email{
			Url:         "golang.org",
			CreatedTime: 1,
		},
	}

	fmt.Println(emp, emp.Addr.City, emp.Url, emp.name)
	// 嵌套结构体同名字段就近

	emp.Get()
	emp.Email.Get()
}
