// kafka生产者
package main

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
)

type Tweet struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 等待ack策略 0 1 -1
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 分配到哪个分区的算法
	config.Producer.Return.Successes = true
	config.ChannelBufferSize = 1024
	config.Net.MaxOpenRequests = 1
	config.Producer.Idempotent = true // 保证消息的幂等性

	tweet := Tweet{
		Username: "karina",
		Message:  "hello i am karina",
	}

	msgData, _ := json.Marshal(tweet)
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test"
	//msg.Key = sarama.StringEncoder("test_key")
	msg.Value = sarama.StringEncoder(msgData)

	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	// 异步发送客户端 发送结果通过channel获取
	aclient, err := sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, config)

	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()
	defer aclient.Close()

	input := aclient.Input()

	input <- msg
	success := aclient.Successes()
	data := <-success
	fmt.Println(data.Topic)

	//pid, offset, err := client.SendMessage(msg)
	//if err != nil {
	//	fmt.Println("send message failed, err: ", err)
	//	return
	//}
	//fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
