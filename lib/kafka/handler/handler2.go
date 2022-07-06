package handler

import (
	"fmt"

	"github.com/Shopify/sarama"
)

type MyHandler struct {
}

func (e *MyHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (e *MyHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (e *MyHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		go func(msg *sarama.ConsumerMessage) {
			fmt.Println("i am handler2 with partition:", msg.Partition)
			session.MarkMessage(msg, "")
		}(msg)
	}
	return nil
}
