package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

type Tweet struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var (
	client *elastic.Client
)

func InitEs()  {
	var err error
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200/"))
	if err != nil {
		log.Fatal("connect es fail: ", err)
	}
	fmt.Println("connect es success")
}

func Insert()  {
	tweet := Tweet{
		Username: "karina",
		Message:  "歌手喜欢唱歌",
	}

	// add index
	_, err := client.Index().
		Index("messages").
		//Id("").
		BodyJson(tweet).
		Do(context.Background())

	if err != nil {
		log.Fatal("add index fail: ", err)
	}

	fmt.Println("add success!")
}

