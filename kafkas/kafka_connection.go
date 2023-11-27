package kafkas

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func SetupKafka() {
	log.Println("Connecting to kafka broker")
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "test", 0)
	_ = conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	_, _ = conn.WriteMessages(kafka.Message{
		Value: []byte("hello kafka again"),
	})
}
