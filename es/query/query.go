package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"reflect"
	"strconv"
	"time"
)

var client *elastic.Client

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func init() {
	var err error
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200/"))
	if err != nil {
		log.Println("connect to es failed:", err)
		return
	}
	fmt.Println("connect to es success")
}

func Add(name string) {
	pro := Product{
		ID:   int(time.Now().Unix()),
		Name: name,
	}

	// 新增 相同id覆盖原记录
	//pro.ID = 1636472116

	_, err := client.Index().
		Index("product").
		Id(strconv.Itoa(pro.ID)).
		BodyJson(pro).
		Do(context.Background())

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("add success.")
	time.Sleep(1 * time.Second)
}

func Get(query string) {
	q := elastic.NewQueryStringQuery("name:" + query)
	res, err := client.Search("product").Query(q).Do(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	var p Product
	for _, v := range res.Each(reflect.TypeOf(p)) {
		if pro, ok := v.(Product); ok {
			fmt.Printf("%#v\n", pro)
		}
	}
}

func Update(id, name string) {
	pid, _ := strconv.Atoi(id)
	pro := Product{
		ID:   pid,
		Name: name,
	}
	_, err := client.Update().Index("product").Id(id).Doc(pro).Do(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("update success.")
}

func Query(q string) {
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("name", q))
	// 范围查询
	boolQ.Filter(elastic.NewRangeQuery("id").Gt(1636534343))
	res, err := client.Search("product").Query(boolQ).Do(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	var p Product
	for _, v := range res.Each(reflect.TypeOf(p)) {
		if pro, ok := v.(Product); ok {
			fmt.Printf("%#v\n", pro)
		}
	}
}

func main() {
	//Add("戴尔笔记本电脑")
	//Add("台式电脑")
	//Add("苹果笔记本电脑")
	//Add("苹果iPhone 大容量电池")
	//Add("电视机")
	//Add("戴尔显示器")

	//Get("电脑")
	Query("电脑")
	//Update("1636470814", "phone")

}
