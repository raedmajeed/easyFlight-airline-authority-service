package config

import (
	"context"
	"github.com/raedmajeed/admin-servcie/pkg/service/interfaces"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaReader struct {
	SearchReader       *kafka.Reader
	SearchSelectReader *kafka.Reader
	svc                interfaces.AdminAirlineService
}

func NewKafkaReaderConnect(svc interfaces.AdminAirlineService) *KafkaReader {
	// trying
	searchReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "search-flight-request-3",
		GroupID: "search-request-3",
	})
	searchSelectReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "selected-flight-request-4",
		GroupID: "search-selected-3",
	})
	return &KafkaReader{
		SearchReader:       searchReader,
		svc:                svc,
		SearchSelectReader: searchSelectReader,
	}
}

type SearchDetails struct {
	DepartureAirport    string
	ArrivalAirport      string
	DepartureDate       string
	ReturnDepartureDate string
	ReturnFlight        bool
	MaxStops            string
}

func (k *KafkaReader) SearchFlightRead(ctx context.Context) {
	for {
		message, _ := k.SearchReader.ReadMessage(ctx)
		select {
		case <-ctx.Done():
			log.Println("context cancelled, terminating")
			return
		default:
			k.svc.SearchFlightInitial(message)
			//err := k.SearchReader.CommitMessages(ctx, message)
			//if err != nil {
			//	return
			//}
		}
	}
}

func (k *KafkaReader) SearchSelectFlightRead(ctx context.Context) {
	for {
		message, _ := k.SearchSelectReader.ReadMessage(ctx)
		select {
		case <-ctx.Done():
			log.Println("context cancelled, terminating")
			return
		default:
			log.Println("message reached in SearchSelectFlightRead() - kafka_reader")
			//break
			k.svc.SearchSelectFlight(ctx, message)
			//err := k.SearchSelectReader.CommitMessages(ctx, message)
			//if err != nil {
			//	return
			//}
			//return
		}
	}
}
