package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Student struct {
	Name string `json:"name"`
}

func main() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("go").Collection("student")

	//s1 := Student{Name: "karina"}
	//result, err := collection.InsertOne(context.Background(), s1)
	//if err != nil {
	//	log.Fatal(err)
	//} else {
	//	log.Println("insert success. ", result.InsertedID)
	//}
	//filter := bson.D{{"name", "karina"}}
	oid, _ := primitive.ObjectIDFromHex("6250327e606db6120a6ba373")
	filter := bson.M{"_id": oid}
	update := bson.D{
		{"$set", bson.D{
			{"name", "my"},
		}}}
	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.ModifiedCount)
}
