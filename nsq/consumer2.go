package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

type Consumer struct {
}

func (c *Consumer) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}

	fmt.Println("get data from nsq: ", string(m.Body), "time: ", time.Unix(m.Timestamp, 0).Format("2006-01-02 15:04:05"))
	return nil
}

// 初始化消费者
func initConsumer(topic, channel, addr string) error {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = 30 * time.Second
	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		return err
	}

	// 消费时回调对象的方法
	consumer := &Consumer{}
	c.AddHandler(consumer)

	// 连接nsq lookupd http
	err = c.ConnectToNSQLookupd(addr)
	if err != nil {
		return err
	}

	//c.Stop()
	return nil
}

func main() {
	err := initConsumer("order_info", "goshop2", "127.0.0.1:4161")
	if err != nil {
		log.Fatal("init consumer failed: " + err.Error())
	}

	for {

	}

}
