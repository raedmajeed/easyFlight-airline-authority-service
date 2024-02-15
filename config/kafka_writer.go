package config

import (
	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	EmailWriter *kafka.Writer
}

func NewKafkaWriterConnect(cfg *ConfigParams) *KafkaWriter {
	emailWriter := kafka.Writer{
		Addr:                   kafka.TCP(cfg.KAFKABROKER),
		Topic:                  "email-service-2",
		AllowAutoTopicCreation: true,
	}

	return &KafkaWriter{
		EmailWriter: &emailWriter,
	}
}
