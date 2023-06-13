package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type Producer struct {
	topic    string
	producer sarama.AsyncProducer
}

func NewProducer(brokers []string, topic string, options ...func(*sarama.Config)) (*Producer, error) {
	kafkaConfig := sarama.NewConfig()
	for _, option := range options {
		option(kafkaConfig)
	}

	producer, err := sarama.NewAsyncProducer(brokers, kafkaConfig)
	if err != nil {
		return nil, err
	}

	go func() {
		for err := range producer.Errors() {
			logrus.Error(fmt.Sprintf("Failed to send log entry to kafka : error: [%v]\n", err.Error()))
			logrus.Error(fmt.Sprintf("Failed to send log entry to kafka : brokers: [%v]\n", brokers))
			logrus.Error(fmt.Sprintf("Failed to send log entry to kafka : kafkaConfig: [%v]\n", kafkaConfig))
			logrus.Error(fmt.Sprintf("Failed to send log entry to kafka : producer: [%v]\n", producer))
		}
	}()

	return &Producer{
		topic:    topic,
		producer: producer,
	}, nil
}

func (p *Producer) Send(v interface{}) error {
	msg, err := json.Marshal(v)
	if err != nil {
		return err
	}

	if p == nil {
		log.Println("Kafka producer is nil")
		return fmt.Errorf("Kafka producer is nil")
	}
	if p.producer == nil {
		log.Println("Kafka producer is nil")
		return fmt.Errorf("Kafka producer is nil")
	}

	p.producer.Input() <- &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(msg),
	}

	return nil
}
func (p *Producer) SendWithKey(v interface{}, key string) error {
	msg, err := json.Marshal(v)
	if err != nil {
		return err
	}

	if p.producer == nil {
		log.Println("Kafka producer is nil")
		return fmt.Errorf("Kafka producer is nil")
	}

	if key == "" {
		log.Println("producer Key is empty")
		return fmt.Errorf("producer Key is empty")
	}

	p.producer.Input() <- &sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(msg),
	}

	return nil
}

func (p *Producer) Close() error {
	return p.producer.Close()
}
