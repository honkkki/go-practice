package main

import (
	"encoding/json"
	"fmt"
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
}
