package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"sync"
)

var Service = newKafkaService([]string{"0.0.0.0:29092"})

// kafkaService struct
type kafkaService struct {
	Brokers   []string
	Producer  sarama.SyncProducer
	Consumer  sarama.Consumer
	Topic     string
	waitGroup sync.WaitGroup
}

// newKafkaService initializes kafkaService
func newKafkaService(brokers []string) *kafkaService {
	//// Producer configuration
	//producerConfig := sarama.NewConfig()
	//producerConfig.Producer.Return.Successes = true
	//
	//producer, err := sarama.NewSyncProducer(brokers, producerConfig)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to create producer: %v", err)
	//}

	// Consumer configuration

	service := &kafkaService{
		Brokers: brokers,
	}

	return service
}

// ProduceMessage sends a message to Kafka
func (k *kafkaService) SendMessage(topic string, message interface{}) error {
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(k.Brokers, producerConfig)
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		return fmt.Errorf("failed to create producer: %v", err)
	}
	defer producer.Close()
	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(jsonData),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}

	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	return nil
}

// ConsumeMessages continuously listens to Kafka topic
func (k *kafkaService) ConsumeMessages(topic string, fn func(pc sarama.PartitionConsumer)) error {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer(k.Brokers, config)
	defer func(consumer sarama.Consumer) {
		if consumer != nil {
			err := consumer.Close()
			if err != nil {
				fmt.Printf("Failed to close consumer: %v", err)
			}
		}

	}(consumer)
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		return err
	}

	for _, partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			return err
		}
		k.waitGroup.Add(1)
		go func(fc func(pc sarama.PartitionConsumer), pc sarama.PartitionConsumer) {
			defer k.waitGroup.Done()
			fc(pc)
		}(fn, pc)
	}
	k.waitGroup.Wait()
	return nil
}
