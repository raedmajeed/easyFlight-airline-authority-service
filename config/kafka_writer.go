package config

import (
	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	EmailWriter        *kafka.Writer
	SearchWriter       *kafka.Writer
	SearchSelectWriter *kafka.Writer
}

func NewKafkaWriterConnect() *KafkaWriter {
	emailWriter := kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "email-service",
	}

	searchWriter := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "search-flight-response-1",
	}
	searchSelectWriter := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "search-flight-response-2",
	}
	return &KafkaWriter{
		EmailWriter:        &emailWriter,
		SearchWriter:       searchWriter,
		SearchSelectWriter: searchSelectWriter,
	}
}
