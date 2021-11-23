package main

import (
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
)

func watch() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal("connect etcd fail: ", err)
	}
	defer cli.Close()

	for {
		ch := cli.Watch(context.Background(), "name")
		for data := range ch {
			for _, v := range data.Events {
				fmt.Printf("watch success, %s - %s : %s\n", v.Type, v.Kv.Key, v.Kv.Value)
			}
		}
	}
}

func main() {
	watch()
}
