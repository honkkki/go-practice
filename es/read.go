package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"reflect"
)

func Read()  {
	//mp := elastic.NewMatchPhraseQuery("message", "i")
	mp := elastic.NewQueryStringQuery("歌手")
	res, err := client.Search("message").Query(mp).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var tweet Tweet
	fmt.Println("search success: ", res.Hits.TotalHits.Value)
	for _, v := range res.Each(reflect.TypeOf(tweet)) {
		t := v.(Tweet)
		fmt.Printf("%#v\n", t)
	}

	fmt.Println("---------------------")
	mp1 := elastic.NewQueryStringQuery("karina")
	res1, err := client.Search("message").Query(mp1).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var tweet1 Tweet
	fmt.Println("search success: ", res1.Hits.TotalHits.Value)
	for _, v1 := range res1.Each(reflect.TypeOf(tweet1)) {
		t1 := v1.(Tweet)
		fmt.Printf("%#v\n", t1)
	}
}

func ReadData(params string)  {
	mp := elastic.NewQueryStringQuery(params)
	res, err := client.Search("message").Query(mp).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var tweet Tweet
	fmt.Println("search success num: ", res.Hits.TotalHits.Value)
	for _, v := range res.Each(reflect.TypeOf(tweet)) {
		t := v.(Tweet)
		fmt.Printf("%#v\n", t)
	}
}




