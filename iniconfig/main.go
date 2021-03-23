package main

import (
	"go-practice/iniconfig/ini"
	"io/ioutil"
)

func main()  {
	data, _ := ioutil.ReadFile("./ini/config.ini")

	conf := &ini.Config{}

	ini.UnMarshal(data, conf)
}
