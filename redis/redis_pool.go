package main

import (
	"fmt"
	"go-practice/iniconfig/ini"
	"io/ioutil"
	"reflect"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	Pool *redis.Pool
)

func init() {
	Pool = newPool()
}

func newPool() *redis.Pool {
	data, _ := ioutil.ReadFile("../config.ini")
	cfg := &ini.Config{}

	ini.UnMarshal(data, cfg)

	host := cfg.RedisConfig.Host
	port := cfg.RedisConfig.Port
	host = "127.0.0.1"
	port = 6379
	addr := fmt.Sprintf("%s:%d", host, port)

	return &redis.Pool{
		MaxIdle:     60,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < 5*time.Minute {
				return nil
			}

			_, err := c.Do("ping")
			return err
		},
	}
}

func main() {
	c := Pool.Get()
	defer c.Close()
	res, _ := redis.String(c.Do("get", "name"))
	fmt.Println(res)

	num, _ := redis.String(c.Do("get", "num"))
	fmt.Println(reflect.TypeOf(num))

	num1, _ := redis.Int(c.Do("get", "num"))
	fmt.Println(reflect.TypeOf(num1))
}
