package main

import (
	"fmt"
	"go-practice/iniconfig/ini"
	"io/ioutil"
	"log"

	"github.com/gomodule/redigo/redis"
)

var (
	r redis.Conn
)

func initRedis() (err error) {
	data, _ := ioutil.ReadFile("../config.ini")
	cfg := &ini.Config{}

	err = ini.UnMarshal(data, cfg)
	if err != nil {
		return err
	}

	host := cfg.RedisConfig.Host
	port := cfg.RedisConfig.Port
	r, err = redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := initRedis()
	if err != nil {
		log.Fatal(err)
	}

	defer r.Close()
	res, err := r.Do("set", "name", "jisoo")
	if err != nil {
		log.Fatal(err)
	}

	res, _ = redis.String(r.Do("get", "name"))
	fmt.Println(res)
}
