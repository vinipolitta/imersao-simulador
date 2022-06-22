package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

func (k *KafkaConsumer) Consumer() {

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatal("Error consuming kafka message:" + err.Error())
	}
	topics := []string{os.Getenv("KafkaReadTopic")}

	c.SubscribeTopics(topics, nil)
}
