package main

import (
	"fmt"
	"go-practice/iniconfig/ini"
	"io/ioutil"
	"log"
)

func main()  {
	data, _ := ioutil.ReadFile("../config.ini")
	cfg := &ini.Config{}

	err := ini.UnMarshal(data, cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.RedisConfig.Host)
}
