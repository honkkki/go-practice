// kafka生产者
package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
)

type Tweet struct {
	Username string `json:"username"`
	Message string `json:"message"`
}

func main()  {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	tweet := Tweet{
		Username: "karina",
		Message:  "hello i am karina",
	}

	msgData, _ := json.Marshal(tweet)
	msg := &sarama.ProducerMessage{}
	msg.Topic = "message"
	//msg.Key = sarama.StringEncoder("test_key")
	msg.Value = sarama.StringEncoder(msgData)

	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed, err: ", err)
		return
	}

	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
