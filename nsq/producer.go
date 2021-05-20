package main

import (
	"bufio"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"strings"
)

var producer *nsq.Producer

func initProducer(addr string) error {
	var err error
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(addr, config)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	addr := "127.0.0.1:4150"
	err := initProducer(addr)
	if err != nil {
		log.Fatal("init producer error: " + err.Error())
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		data, _ := reader.ReadString('\n')
		data = strings.TrimSpace(data)
		if data == "stop" {
			break
		}

		err = producer.Publish("order_info", []byte(data))
		if err != nil {
			fmt.Println("publish failed ", err)
			continue
		}

		fmt.Println("publish success!")
	}

}
