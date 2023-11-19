package service

import (
	"context"
	"encoding/json"
	"fmt"
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaPath struct {
	DirectPath []dom.Path
	ReturnPath []dom.Path
}

func (svc *AdminAirlineServiceStruct) SearchFlightInitial(message kafka.Message) {
	flightPath, returnFlightPath, err := svc.SearchFlight(message)
	fmt.Println()
	marshal, err := json.Marshal(KafkaPath{
		DirectPath: flightPath,
		ReturnPath: returnFlightPath,
	})
	if err != nil {
		return
	}

	err = svc.kfk.SearchWriter.WriteMessages(
		context.Background(), kafka.Message{
			Value: marshal,
		})
	if err != nil {
		log.Println("error writing to kafka, error: ", err)
		return
	}
}
