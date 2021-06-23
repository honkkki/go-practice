package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("Failed to start consumer: %s\n", err)
		return
	}
	partitionList, err := consumer.Partitions("test_kafka")
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("test_kafka", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(pc1 sarama.PartitionConsumer) {
			for msg := range pc1.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}


	for  {

	}
	//consumer.Close()

}
