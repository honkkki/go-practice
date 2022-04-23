package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

var (
	client *mongo.Client
)

func init() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

func insert() {
	clt := client.Database("go").Collection("user")
	s1 := Student{Name: "sona", Age: 20}
	result, err := clt.InsertOne(context.Background(), s1)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("insert success. ", result.InsertedID)
	}
}

func update() {
	collection := client.Database("go").Collection("user")
	oid, _ := primitive.ObjectIDFromHex("")
	filter := bson.M{"_id": oid}
	update := bson.D{
		{"$set", bson.D{
			{"age", 25},
		}},
	}
	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.ModifiedCount)
}

func query() {
	clt := client.Database("go").Collection("user")
	opts := options.Find().SetSort(bson.D{{"age", 1}})
	c, err := clt.Find(context.Background(), bson.D{{"name", "karina"}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	var res []bson.M
	if err := c.All(context.Background(), &res); err != nil {
		log.Fatal(err)
	}

	for _, v := range res {
		fmt.Println(v)
	}
}

func main() {
	//insert()
	query()
}
