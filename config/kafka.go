package config

import "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaConfig struct {
	brokerList                string
	topic					  string
	acks                      int
	maxQueueSize              int
}

func (kc KafkaConfig) BrokerList() string {
	return kc.brokerList
}

func (kc KafkaConfig) Topic() string {
	return kc.topic
}

func (kc KafkaConfig) Acks() int {
	return kc.acks
}

func (kc KafkaConfig) MaxQueueSize() int {
	return  kc.maxQueueSize
}

func NewKafkaConfig() KafkaConfig {
	kc := KafkaConfig{
		brokerList:                mustGetString("KAFKA_BROKER_LIST"),
		topic:					   mustGetString("KAFKA_TOPIC"),
		acks:                      mustGetInt("KAFKA_ACKS"),
		maxQueueSize:              mustGetInt("KAFKA_QUEUE_SIZE"),
	}
	return kc
}

func (cfg KafkaConfig) ToKafkaConfigMap() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers":            cfg.BrokerList(),
		"acks":                         cfg.Acks(),
	}
}
