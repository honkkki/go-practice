package main

import "fmt"

type User struct {
	name string
}

func main() {
	m := make(map[string]User)
	user, ok := m["name"]
	if !ok {
		user.name = "tom"
	}

	m["user"] = user
	fmt.Println(m["user"])
}
