package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type Game struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price *float64 `json:"price,omitempty"` // 空值json不处理
}

func main() {
	price := 0.0
	game := Game{
		ID:    1,
		Name:  "rainbow6",
		Price: &price,
	}

	b, _ := jsoniter.Marshal(game)
	fmt.Println(string(b))

	game2 := Game{
		ID:   2,
		Name: "lol",
	}
	b2, _ := jsoniter.Marshal(game2)
	fmt.Println(string(b2))

	b3, _ := jsoniter.MarshalIndent(game, "", "  ")
	fmt.Println(string(b3))
}
