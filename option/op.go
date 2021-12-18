package main

import "fmt"

type Conn struct {
	id      int
	timeout int
	name    string
}

type Option func(conn *Conn)

func WithName(name string) Option {
	return func(conn *Conn) {
		conn.name = name
	}
}

func WithTimeout(timeout int) Option {
	return func(conn *Conn) {
		conn.timeout = timeout
	}
}

func NewConn(id int, option ...Option) *Conn {
	conn := &Conn{id: id}
	for _, op := range option {
		op(conn)
	}

	return conn
}

func main() {
	c := NewConn(1, WithName("tcp"), WithTimeout(10))
	fmt.Printf("%+v\n", c)
}
