package main

import (
	"ats/repo"
	"ats/service"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/util/gconv"
	"log"
	"strings"
	"time"

	cfs "ats/config"
)

func main() {
	service.InitCon()

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer([]string{cfs.BrokerList}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(cfs.Topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start Kafka partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			data := string(msg.Value)
			fmt.Printf("Received message: %s\n", data)
			go simulateSendMessage(data)
		case err := <-partitionConsumer.Errors():
			log.Printf("Error: %v", err)
		}
	}
}

func simulateSendMessage(data string) {
	parts := strings.Split(data, "|")
	if len(parts) < 3 {
		log.Printf("Invalid message format: %s", data)
		return
	}

	activityID := parts[0]
	phone := parts[1]
	message := parts[2]

	time.Sleep(2 * time.Second)
	fmt.Printf("Message sent for activity %s to %s: %s\n", activityID, phone, message)
	//TODO 根据发送消息状态写入消息表
	var ret = "success"
	msg := repo.Messages{
		ActivityId: gconv.Int32(activityID),
		Phone:      phone,
		Message:    message,
		Status:     ret,
		SendTime:   time.Now().Format(time.DateTime),
	}
	err := service.InsertMessage(msg)
	fmt.Println("err: ", err)
}
