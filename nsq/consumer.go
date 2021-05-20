package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

type Consumer struct {
}

func (c *Consumer) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}

	fmt.Println("get data from nsq: ", string(m.Body))
	return nil
}

// 初始化消费者
func initConsumer(topic, channel, addr string) error {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = 10 * time.Second
	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		return err
	}

	consumer := &Consumer{}
	c.AddHandler(consumer)

	err = c.ConnectToNSQLookupd(addr)
	if err != nil {
		return err
	}

	//c.Stop()
	return nil
}

func main() {
	err := initConsumer("order_info", "goshop", "127.0.0.1:4161")
	if err != nil {
		log.Fatal("init consumer failed: " + err.Error())
	}

	for {

	}

}
