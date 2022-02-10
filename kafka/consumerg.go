package main

import (
	"context"
	"go-practice/kafka/handler"
	"log"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	cfg *sarama.Config
)

func init() {
	cfg = sarama.NewConfig()
	cfg.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	cfg.Consumer.Offsets.AutoCommit.Enable = true
	cfg.Consumer.Offsets.Initial = sarama.OffsetNewest
}

func NewConsumerGroup(group string) sarama.ConsumerGroup {
	cg, err := sarama.NewConsumerGroup([]string{"127.0.0.1:9092"}, group, cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cg
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	cg := NewConsumerGroup("group1")
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
			err := cg.Consume(ctx, []string{"test"}, &handler.EventHandler{})
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
