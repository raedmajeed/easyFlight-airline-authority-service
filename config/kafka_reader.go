package config

import (
	"context"
	"fmt"
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
		Topic:   "search-flight-request",
		GroupID: "search-request-1",
	})
	searchSelectReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "search-select-request",
		GroupID:  "search-select",
		MaxBytes: 10e5,
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
	//messageChan := make(chan kafka.Message)
	for {
		message, _ := k.SearchReader.FetchMessage(ctx)
		select {
		case <-ctx.Done():
			log.Println("context cancelled, terminating")
			return
		//case messageChan <- message:
		default:
			fmt.Println(message.Value)
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

func (k *KafkaReader) SearchSelectFlightRead(ctx context.Context) {
	//messageChan := make(chan kafka.Message)
	// here when searchFlight dosen't do anything it gets stuck handle that
	//newCont, cancel := context.WithTimeout(ctx, time.Second*20)
	//defer cancel()
	for {
		message, _ := k.SearchSelectReader.FetchMessage(ctx)
		select {
		case <-ctx.Done():
			log.Println("context cancelled, terminating")
			return
		//case messageChan <- message:
		default:
			fmt.Println(message.Key)
			log.Println("message reached in SearchSelectFlightRead() - kafka_reader")
			//break
			k.svc.SearchSelectFlight(ctx, message)
			err := k.SearchSelectReader.CommitMessages(ctx, message)
			if err != nil {
				return
			}
			return
		}
		//break
	}
	//return messageChan
}
