package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	m := make(map[string]string)

	str := `
{
	"name": "karina"
}
`

	json.Unmarshal([]byte(str), &m)
	fmt.Println(m)

	x := 1
	{
		x = 2
		x := 2
		fmt.Println(x)
		x = 3
	}

	fmt.Println(x)

	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	// io stream decode
	json.NewDecoder(resp.Body)
}
