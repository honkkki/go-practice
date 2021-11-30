package main

import (
	"encoding/json"
	"fmt"
)

type Game struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Desc  string `json:"desc,omitempty"`
}

func main() {
	var res []Game
	g1 := Game{
		Name:  "rainbow",
		Price: 100,
		Desc:  "rainbow",
	}

	g2 := Game{
		Name:  "cs",
		Price: 50,
		Desc:  "",
	}

	res = append(res, g1, g2)
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
}
