package main

import "log"

type Piface interface {
	Name() string
}

type Siface interface {
	//Piface

	IP() string
}

type someObj struct {
	name string
	ip   string
}

func (s *someObj) Name() string {
	return s.name
}

func (s *someObj) IP() string {
	return s.ip
}

func NewObj() Piface {
	return &someObj{
		name: "obj",
		ip:   "0.0.0.0",
	}
}

func main() {
	s := NewObj()
	log.Println(s.(Siface).IP())
}
