package config

import (
	"context"
	"github.com/raedmajeed/admin-servcie/pkg/service/interfaces"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaReader struct {
	SearchReader *kafka.Reader
	svc          interfaces.AdminAirlineService
}

func NewKafkaReaderConnect(svc interfaces.AdminAirlineService) *KafkaReader {
	searchReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "search-flight-request",
		GroupID: "search-request-1",
	})
	return &KafkaReader{
		SearchReader: searchReader,
		svc:          svc,
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
	//messageChan := make(chan kafka.Message)
	for {
		message, _ := k.SearchReader.FetchMessage(ctx)
		select {
		case <-ctx.Done():
			log.Println("context cancelled, terminating")
			return
		//case messageChan <- message:
		default:
			k.svc.SearchFlightInitial(message)
			err := k.SearchReader.CommitMessages(ctx, message)
			if err != nil {
				return
			}
			//return
		}
	}
	//return messageChan
}
