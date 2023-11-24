package service

import (
	"context"
	"encoding/json"
	"fmt"
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	"github.com/segmentio/kafka-go"
	"log"
	"strconv"
)

func (svc *AdminAirlineServiceStruct) SearchFlightInitial(message kafka.Message) {
	flightPath, returnFlightPath, err := svc.SearchFlight(message)
	marshal, err := json.Marshal(dom.KafkaPath{
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

func (svc *AdminAirlineServiceStruct) SearchSelectFlight(ctx context.Context, message kafka.Message) {
	// this is where the logic happens. what is that i should do if error or something happens
	var kafkaPath dom.KafkaPath
	//var selectReq dom.SelectRequest
	var directPath dom.Path
	var returnPath dom.Path

	fmt.Println(string(message.Value)) // printed

	selectReq := dom.SelectRequest{}
	err := json.Unmarshal(message.Value, &selectReq)
	if err != nil {
		log.Println("error marshaling json SearchSelectFlight() - booking service 1")
		//err := svc.kfk.SearchSelectWriter.WriteMessages(ctx, kafka.Message{})
		//writeToKafkaOnError(ctx, svc, &dom.CompleteFlightFacilities{})
		return
	}

	directPathId, _ := strconv.Atoi(selectReq.DirectPathId)
	returnPathID, err := strconv.Atoi(selectReq.ReturnPathId)

	if err != nil || selectReq.ReturnPathId == "" {
		log.Printf("return path id is null so setting it with -1, returnpathid: %v", selectReq.ReturnPathId)
		returnPathID = -1
	}

	// this gets the additional details stored in token, like no of adults, children and cabin class
	//addDetails := selectReq.AddDetails
	token := selectReq.Token

	redisVal := svc.redis.Get(ctx, token)
	if err = json.Unmarshal([]byte(redisVal.Val()), &kafkaPath); err != nil {
		fmt.Println("error unmarshalling json from redis SearchSelectFlight() - booking service")
		// add retry and if retry fails still write to kafka again
		//writeToKafkaOnError(ctx, svc, &dom.CompleteFlightFacilities{})
		return
	}

	fmt.Println(kafkaPath.DirectPath, "redis path")

	cabinClass := "ECONOMY"
	if !selectReq.Economy {
		cabinClass = "BUSINESS"
	}
	log.Printf("cabin class is %v", cabinClass)

	var returnPathList []dom.Path
	if returnPathID != -1 {
		returnPathList = kafkaPath.ReturnPath
	}
	directPathList := kafkaPath.DirectPath

	for _, p := range directPathList {
		if p.PathId == directPathId {
			directPath = p
			break
		}
	}

	for _, p := range returnPathList {
		if p.PathId == returnPathID {
			returnPath = p
			break
		}
	}

	if directPath.PathId <= 0 {
		log.Println("path id not present, chance is paths from redis has been deleted")
		// write to kafka here
		//writeToKafkaOnError(ctx, svc, &dom.CompleteFlightFacilities{})
		return
	}

	// this collects the cancellation, baggage and seat availability of direct path flights
	check, directFlight := GetFlightFacilities(svc, directPath, selectReq)
	if !check {
		cf := dom.CompleteFlightFacilities{
			DirectFlight: dom.FlightFacilities{},
			ReturnFlight: dom.FlightFacilities{},
		}
		writeToKafka(ctx, svc, &cf)
		return
	}

	if returnPathID == -1 {
		cf := dom.CompleteFlightFacilities{
			DirectFlight:     directFlight,
			ReturnFlight:     dom.FlightFacilities{},
			NumberOfAdults:   selectReq.Adults,
			NumberOfChildren: selectReq.Children,
			CabinClass:       cabinClass,
		}
		writeToKafka(ctx, svc, &cf)
		return
	}

	check, returnFlight := GetFlightFacilities(svc, returnPath, selectReq)
	if !check {
		log.Println("unable to fetch the return flight details SearchSelectFlight() - kafka_services")
		cf := dom.CompleteFlightFacilities{
			DirectFlight:     directFlight,
			ReturnFlight:     dom.FlightFacilities{},
			NumberOfAdults:   selectReq.Adults,
			NumberOfChildren: selectReq.Children,
			CabinClass:       cabinClass,
		}
		writeToKafka(ctx, svc, &cf)
		return
	}
	cf := dom.CompleteFlightFacilities{
		DirectFlight:     directFlight,
		ReturnFlight:     returnFlight,
		NumberOfAdults:   selectReq.Adults,
		NumberOfChildren: selectReq.Children,
		CabinClass:       cabinClass,
	}
	writeToKafka(ctx, svc, &cf)
	return
}

// this method is invoked when some error happens
func writeToKafkaOnError(ctx context.Context, svc *AdminAirlineServiceStruct, cf *dom.CompleteFlightFacilities) {
	marshal, _ := json.Marshal(&cf)
	// implement retry mechanism here
	err := svc.kfk.SearchSelectWriter.WriteMessages(ctx, kafka.Message{
		Value: marshal,
	})
	if err != nil {
		// write to kafka here
		log.Println("error writing to kafka writeToKafkaOnError() - kafka_services 2")
		return
	}
}

func writeToKafka(ctx context.Context, svc *AdminAirlineServiceStruct, cf *dom.CompleteFlightFacilities) {
	marshal, err := json.Marshal(&cf)
	if err != nil {
		log.Println("error marshaling json writeToKafka() - kafka_services 1")
		//writeToKafkaOnError(ctx, svc, &dom.CompleteFlightFacilities{})
		return
	}
	// implement retry mechanism here
	err = svc.kfk.SearchSelectWriter.WriteMessages(ctx, kafka.Message{
		Value: marshal,
	})
	if err != nil {
		// write to kafka here
		log.Println("error writing to kafka writeToKafka() - kafka_services 1")
		//writeToKafkaOnError(ctx, svc, &dom.CompleteFlightFacilities{})
		return
	}
}

func GetFlightFacilities(svc *AdminAirlineServiceStruct, path dom.Path, addDetails dom.SelectRequest) (bool, dom.FlightFacilities) {
	flightNumber := path.Flights[0].FlightNumber
	flight, err := svc.repo.FindFlightByFlightNumber(flightNumber)
	if err != nil {
		log.Println("unable to get flight details GetFlightFacilities() - kafka_services")
		return false, dom.FlightFacilities{}
	}

	cancellation, err := getCancellationDetails(svc, *flight)
	baggage, err := getBaggageDetails(svc, *flight)
	seatCheckD, _ := seatAvailabilityCheck(svc, path.Flights, addDetails)
	//seatCheckR, ReturnSeats := seatAvailabilityCheck(svc, path.Flights, addDetails)
	if !seatCheckD {
		return false, dom.FlightFacilities{}
	}

	//bookSeats(DirectSeats, addDetails, svc, flight.ID)
	canc := dom.Cancellation{
		CancellationPercentage:     cancellation.CancellationPercentage,
		CancellationDeadlineBefore: cancellation.CancellationDeadlineBefore,
		Refundable:                 cancellation.Refundable,
	}

	bag := dom.Baggage{
		CabinAllowedHeight:  baggage.CabinAllowedHeight,
		CabinAllowedWeight:  baggage.CabinAllowedWeight,
		CabinAllowedBreadth: baggage.CabinAllowedBreadth,
		CabinAllowedLength:  baggage.CabinAllowedLength,
		HandAllowedHeight:   baggage.HandAllowedHeight,
		HandAllowedWeight:   baggage.HandAllowedWeight,
		HandAllowedBreadth:  baggage.HandAllowedBreadth,
		HandAllowedLength:   baggage.HandAllowedLength,
		FeeExtraPerKGHand:   baggage.FeeExtraPerKGHand,
		FeeExtraPerKGCabin:  baggage.FeeExtraPerKGCabin,
		Restrictions:        baggage.Restrictions,
	}
	return true, dom.FlightFacilities{
		Cancellation: canc,
		Baggage:      bag,
		FlightPath:   path,
	}
}

func getCancellationDetails(svc *AdminAirlineServiceStruct, flight dom.FlightFleets) (dom.AirlineCancellation, error) {
	baggageId := flight.BaggagePolicyID
	baggage, err := svc.repo.FindAirlineCancellationByid(int32(baggageId))
	return *baggage, err
}

func getBaggageDetails(svc *AdminAirlineServiceStruct, flight dom.FlightFleets) (dom.AirlineBaggage, error) {
	cancellationId := flight.CancellationPolicyID
	cancellation, err := svc.repo.FindAirlineBaggageByid(int32(cancellationId))
	return *cancellation, err
}

func seatAvailabilityCheck(svc *AdminAirlineServiceStruct, flights []dom.FlightDetails, addDetails dom.SelectRequest) (bool, []dom.TemporaryData) {
	var tempData []dom.TemporaryData
	for _, f := range flights {
		chartID := f.FlightChartID
		seat, _ := svc.repo.FindSeatsByChartID(chartID)
		if addDetails.Economy {
			if (addDetails.Adults + addDetails.Children) > seat.EconomySeatTotal-seat.EconomySeatBooked {
				log.Printf("economy seats in flight %v overbooked", f.FlightNumber)
				return false, nil
			}
		} else {
			log.Printf("business seats in flight %v overbooked", f.FlightNumber)
			if (addDetails.Adults + addDetails.Children) > seat.BusinessSeatTotal-seat.BusinessSeatBooked {
				return false, nil
			}
		}
		tempData = append(tempData, dom.TemporaryData{
			ChartId: chartID,
			Seats:   seat,
		})
	}
	return true, tempData
}

/*func bookSeats(seats []TemporaryData, add SearchClaims, svc *AdminAirlineServiceStruct, id uint) {
	for _, data := range seats {
		cId := data.chartId
		seat := data.seats
		if add.Economy {
			seat.EconomySeatBooked = seat.EconomySeatBooked + add.Adults + add.Children
			svc.repo.UpdateBookedSeat(seat, cId)
		} else {
			seat.BusinessSeatBooked = seat.BusinessSeatBooked + add.Adults + add.Children
			svc.repo.UpdateBookedSeat(seat, cId)
		}
	}
}*/
