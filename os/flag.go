package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	length  int
	charset string
)

func GetChars() string {
	a := 'a'
	s1 := make([]byte, 0, 26)
	//fmt.Printf("%T \n", a)	// int32
	for i := 0; i < 26; i++ {
		s1 = append(s1, byte(a))
		a++
	}

	A := 'A'
	s2 := make([]byte, 0, 26)
	for i := 0; i < 26; i++ {
		s2 = append(s2, byte(A))
		A++
	}

	num := make([]string, 0, 10)

	for i := 0; i < 10; i++ {
		num = append(num, strconv.Itoa(i))
	}

	s := append(s1, s2...)
	fmt.Println(num)
	return string(s) + strings.Join(num, "")
}

// 获取命令行参数
func getArgs() {
	flag.IntVar(&length, "l", 8, "length of password")
	flag.StringVar(&charset, "c", "char", "charset of password")
	flag.Parse()
}

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(10))
	chars := GetChars()
	fmt.Println(chars)
	getArgs()
	fmt.Printf("%d \n", length)
}
