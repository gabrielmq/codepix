package kafka

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
	}

	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}

	return p
}

func Publish(msg, topic string, producer *ckafka.Producer, deliveryChan chan ckafka.Event) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}

	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}

	return nil
}

func DeliveryReport(deliveryChan chan ckafka.Event) {
	for delivery := range deliveryChan {
		switch event := delivery.(type) {
		case *ckafka.Message:
			if event.TopicPartition.Error != nil {
				fmt.Println("Delivery failed: ", event.TopicPartition)
			} else {
				fmt.Println("Delivered message to: ", event.TopicPartition)
			}
		}
	}
}
