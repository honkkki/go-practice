package main

import (
	"bytes"
	"fmt"
	"net/url"
)

func main() {
	buf := new(bytes.Buffer)
	buf.WriteString("hello golang")
	fmt.Println(buf.String())
	str := "https://www.google.com/"
	u, _ := url.Parse(str)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
}
