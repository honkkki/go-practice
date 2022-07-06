package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal("connect etcd fail: ", err)
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	_, err = cli.Put(ctx, "name", "hello karina")
	defer cancel()
	if err != nil {
		log.Fatal("put fail: ", err)
	}

	fmt.Println("put to etcd success!")
}
