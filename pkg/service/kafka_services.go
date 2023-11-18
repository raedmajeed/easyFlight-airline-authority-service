package service

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func (svc *AdminAirlineServiceStruct) SearchFlightInitial(message kafka.Message) {
	svc.SearchFlight(message)
	err := svc.kfk.SearchWriter.WriteMessages(
		context.Background(), kafka.Message{
			Value: []byte("maybe good for better is aslam"),
		})
	if err != nil {
		log.Println("error writing to kafka, error: ", err)
		return
	}
}
