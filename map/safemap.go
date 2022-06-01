package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/collection"
)

func main() {
	m := collection.NewSafeMap()
	m.Set("name", "tom")
	fmt.Println(m.Size())
}
