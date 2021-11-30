package main

import "fmt"

type Person struct {
	Age int
}

func main() {
	//m := map[string]Person{
	//	"p": Person{Age: 20},
	//}

	//m["p"].Age = 23		// error 无法对map的值寻址

	m := map[string]string{}
	m["name"] = "jisoo"
	m["name"] = "karina"
	fmt.Println(m)
}
