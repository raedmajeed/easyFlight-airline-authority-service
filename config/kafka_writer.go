package config

import (
	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	EmailWriter        *kafka.Writer
	SearchWriter       *kafka.Writer
	SearchSelectWriter *kafka.Writer
}

func NewKafkaWriterConnect(cfg *ConfigParams) *KafkaWriter {
	emailWriter := kafka.Writer{
		Addr:                   kafka.TCP(cfg.KAFKABROKER),
		Topic:                  "email-service-2",
		AllowAutoTopicCreation: true,
	}

	searchWriter := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "search-flight-response-3",
		AllowAutoTopicCreation: true,
	}
	searchSelectWriter := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "selected-flight-response-2",
		AllowAutoTopicCreation: true,
	}
	return &KafkaWriter{
		EmailWriter:        &emailWriter,
		SearchWriter:       searchWriter,
		SearchSelectWriter: searchSelectWriter,
	}
}
