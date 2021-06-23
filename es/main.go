package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

type Tweet struct {
	Username string
	Message  string
}

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200/"))
	if err != nil {
		log.Fatal("connect es fail: ", err)
	}
	fmt.Println("connect es success")

	tweet := Tweet{
		Username: "karina",
		Message:  "hello karina!",
	}

	_, err = client.Index().
		Index("info").
		Type("tweet").
		Id("1").
		BodyJson(tweet).
		Do(context.Background())

	if err != nil {
		log.Fatal("add index fail: ", err)
	}

	fmt.Println("add success!")
}
