package handler

import (
	"fmt"

	"github.com/Shopify/sarama"
)

type EventHandler struct {
}

func (e *EventHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (e *EventHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (e *EventHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		// consume
		msg := msg
		go func() {
			fmt.Println("i am handler1 with partition:", msg.Partition)
			fmt.Println(string(msg.Value), "Partition", msg.Partition, "offset", msg.Offset)
			// 处理消息成功后标记为处理, 然后会自动提交
			session.MarkMessage(msg, "")
		}()

	}
	return nil
}
