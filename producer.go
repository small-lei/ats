package main

import (
	cfs "ats/config"
	"encoding/csv"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
	"time"

	"ats/service"

	"github.com/gogf/gf/v2/util/gconv"
)

const (
	brokerList = "localhost:9092"
	topic      = "message_topic"
)

func readCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{cfs.BrokerList}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka producer: %v", err)
	}
	defer producer.Close()

	service.InitCon()

	activities, err := readCSV("activities.csv")
	if err != nil {
		log.Fatalf("Failed to read activities CSV: %v", err)
	}

	recipients, err := readCSV("recipients.csv")
	if err != nil {
		log.Fatalf("Failed to read recipients CSV: %v", err)
	}

	for _, activity := range activities[1:] {
		activityID := gconv.Int32(activity[0])
		template := activity[1]
		scheduledTime, _ := time.Parse(time.RFC3339, activity[2])
		// 计算延迟时间
		delay := time.Until(scheduledTime)
		if delay > 0 {
			time.Sleep(delay)
		}

		for _, recipient := range recipients[1:] {
			phone := recipient[0]
			name := recipient[1]
			message := fmt.Sprintf(template, name)
			//检查活动用户是否发送
			senderFlag, _ := service.CheckActUserSender(activityID, phone)
			if senderFlag != nil {
				continue
			}
			msg := &sarama.ProducerMessage{
				Topic: cfs.Topic,
				Value: sarama.StringEncoder(fmt.Sprintf("%s|%s|%s", activityID, phone, message)),
			}

			_, _, err := producer.SendMessage(msg)
			if err != nil {
				log.Printf("Failed to send message to Kafka: %v", err)
			}
		}
	}
}
