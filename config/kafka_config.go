package config

import (
	"github.com/segmentio/kafka-go"
)

type KafkaReadWrite struct {
	EmailWriter *kafka.Writer
}

func NewKafkaConnect() *KafkaReadWrite {
	emailWriter := kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "email-service",
	}
	return &KafkaReadWrite{
		EmailWriter: &emailWriter,
	}
}
