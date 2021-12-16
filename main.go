package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	brokerAddrs := []string{"localhost:9092"}
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	admin, err := sarama.NewClusterAdmin(brokerAddrs, config)
	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
	}
	defer func() { _ = admin.Close() }()
	err = admin.CreateTopic("topic.test.1", &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}, false)
	if err != nil {
		log.Fatal("Error while creating topic: ", err.Error())
	}
}
