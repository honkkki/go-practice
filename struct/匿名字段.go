package main

import (
	"fmt"
	"go-practice/struct/user"
)

type Addr struct {
	Province string
	City string
	CreatedTime int
}

type Email struct {
	Url string
	CreatedTime int
}

type Emp struct {
	name string
	Addr Addr
	Email
}

func main() {
	var u user.User
	u.Name = "jisoo"

	fmt.Println(u)

	emp := Emp{
		name:  "jisoo",
		Addr:  Addr{
			Province:    "gd",
			City:        "sz",
			CreatedTime: 1,
		},
		Email: Email{
			Url:         "qq@qq.com",
			CreatedTime: 1,
		},
	}

	fmt.Println(emp, emp.Addr.City, emp.Url)

}
