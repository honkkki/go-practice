// 消费kafka队列中的数据，插入到elastic中
package main

import (
	"encoding/json"
	"fmt"
	"go-practice/es"
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

func main() {
	es.InitEs()
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("failed to start consumer: %s\n", err)
		return
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	partitionList, err := consumer.Partitions("message")
	if err != nil {
		fmt.Println("failed to get the list of partitions: ", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("message", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(pc1 sarama.PartitionConsumer) {
			for msg := range pc1.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				// 处理消息
				var tweet es.Tweet
				err := json.Unmarshal(msg.Value, &tweet)
				if err != nil {
					fmt.Println("unmarshal fail: ", err)
				}
				fmt.Printf("%#v\n", tweet)
				es.InsertData(tweet)
			}
		}(pc)
	}

	select {
	case <-c:
		fmt.Println("\nclose consumer process")
	}

}
