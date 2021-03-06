package main

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
	"log"
	"time"
)

func conn()  {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Fatal("connect etcd fail: ", err)
	}
	defer cli.Close()
	fmt.Println("connect etcd success!")
}

func put()  {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal("connect etcd fail: ", err)
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	_, err = cli.Put(ctx, "name", "karina")
	cancel()
	if err != nil {
		log.Fatal("put fail: ", err)
	}
	fmt.Println("put to etcd success!")
}

func get()  {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal("connect etcd fail: ", err)
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	resp, err := cli.Get(ctx, "name")
	cancel()
	if err != nil {
		log.Fatal("get fail: ", err)
	}

	for _, v := range resp.Kvs {
		fmt.Printf("get %s : %s\n", v.Key, v.Value)
	}
}

func main()  {
	//conn()
	//put()
	get()
}
