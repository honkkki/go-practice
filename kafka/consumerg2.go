package main

import (
	"context"
	"go-practice/kafka/handler"
	"log"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	cfg2 *sarama.Config
)

func init() {
	cfg2 = sarama.NewConfig()
	cfg2.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	cfg2.Consumer.Offsets.AutoCommit.Enable = true
	cfg2.Consumer.Offsets.Initial = sarama.OffsetNewest
}

func NewConsumerGroup2(group string) sarama.ConsumerGroup {
	cg, err := sarama.NewConsumerGroup([]string{"127.0.0.1:9092"}, group, cfg2)
	if err != nil {
		log.Fatal(err)
	}

	return cg
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	cg := NewConsumerGroup2("group1")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func() {
		if err := cg.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			// rebalance by kafka will return
			err := cg.Consume(ctx, []string{"test"}, &handler.MyHandler{})
			if err != nil {
				log.Fatal(err)
			}

			// maybe cancel by sarama
			if ctx.Err() != nil {
				return
			}
		}
	}()

	wg.Wait()
}
