package service

import (
	"context"
	"encoding/json"
	"errors"
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"
)

func (svc *AdminAirlineServiceStruct) SearchFlightInitial(ctx context.Context, p *pb.SearchFlightRequestAdmin) (*pb.SearchFlightResponseAdmin, error) {
	search := dom.SearchDetails{
		DepartureAirport:    p.DepartureAirport,
		DepartureDate:       p.DepartureDate,
		ArrivalAirport:      p.ArrivalAirport,
		ReturnDepartureDate: p.ReturnDepartureDate,
		ReturnFlight:        p.ReturnFlight,
		Economy:             p.Economy,
		MaxStops:            p.MaxStops,
	}

	flightPath, returnFlightPath, err := svc.SearchFlight(search)
	if err != nil {
		return nil, err
	}

	// write to grpc here
	flightPathResponse := convertPathsToResponse(flightPath)
	returnFlightPathResponse := convertPathsToResponse(returnFlightPath)

	searchResponse := &pb.SearchFlightResponseAdmin{
		DirectPath:       flightPathResponse,
		ReturnPath:       returnFlightPathResponse,
		DepartureAirport: p.DepartureAirport,
		ArrivalAirport:   p.ArrivalAirport,
	}
	return searchResponse, nil
}

func convertPathsToResponse(paths []dom.Path) []*pb.Path {
	var result []*pb.Path
	for _, fl := range paths {
		var flights []*pb.FlightDetailsAdmin
		// converts the flightDetails to pb.FlightDetails
		for _, f := range fl.Flights {
			dt := timestamppb.New(f.DepartureDateTime)
			at := timestamppb.New(f.ArrivalDateTime)
			flights = append(flights, &pb.FlightDetailsAdmin{
				FlightChartId:     uint32(f.FlightChartID),
				FlightNumber:      f.FlightNumber,
				Airline:           f.Airline,
				DepartureAirport:  f.DepartureAirport,
				ArrivalAirport:    f.ArrivalAirport,
				DepartureDate:     f.DepartureDate,
				ArrivalDate:       f.ArrivalDate,
				DepartureTime:     f.DepartureTime,
				ArrivalTime:       f.ArrivalTime,
				DepartureDateTime: dt,
				ArrivalDateTime:   at,
				EconomyFare:       f.EconomyFare,
				BusinessFare:      f.BusinessFare,
			})
		}
		result = append(result, &pb.Path{
			PathId:           int32(fl.PathId),
			Flights:          flights,
			NumberOfStops:    int32(fl.NumberOfStops),
			DepartureAirport: fl.DepartureAirport,
			ArrivalAirport:   fl.ArrivalAirport,
		})
	}
	return result
}

func (svc *AdminAirlineServiceStruct) SearchSelectFlight(ctx context.Context, selectReq *pb.SelectFlightAdmin) (*pb.CompleteFlightDetails, error) {
	directPathId, _ := strconv.Atoi(selectReq.DirectPathId)
	returnPathID, err := strconv.Atoi(selectReq.ReturnPathId)

	if err != nil || selectReq.ReturnPathId == "" {
		log.Printf("return path id is null so setting it with -1, returnpathid: %v", selectReq.ReturnPathId)
		returnPathID = -1
	}
	token := selectReq.Token

	redisVal := svc.redis.Get(ctx, token)
	var searchDetails pb.SearchFlightResponse1
	if err = json.Unmarshal([]byte(redisVal.Val()), &searchDetails); err != nil {
		return nil, err
	}

	cabinClass := "ECONOMY"
	if !selectReq.Economy {
		cabinClass = "BUSINESS"
	}

	var returnPathList []*pb.SearchFlightDetails
	if returnPathID != -1 {
		returnPathList = searchDetails.ReturnFlights
	}
	directPathList := searchDetails.ToFlights

	var directPath *pb.SearchFlightDetails
	for _, p := range directPathList {
		if p.PathId == int32(directPathId) {
			directPath = p
			break
		}
	}

	var returnPath *pb.SearchFlightDetails
	for _, p := range returnPathList {
		if p.PathId == int32(returnPathID) {
			returnPath = p
			break
		}
	}

	if directPath.PathId <= 0 {
		return nil, errors.New("path id not present, chance is paths from redis has been deleted")
	}

	// this collects the cancellation, baggage and seat availability of direct path flights
	check, directFlight := GetFlightFacilities(svc, directPath, selectReq)
	if !check {
		return nil, errors.New("unable to fetch direct flight details")
	}

	if returnPathID == -1 {
		cf := &pb.CompleteFlightDetails{
			DirectFlight:     directFlight,
			ReturnFlight:     nil,
			NumberOfAdults:   selectReq.Adults,
			NumberOfChildren: selectReq.Children,
			CabinClass:       cabinClass,
			DepartureAirport: searchDetails.DepartureAirport,
			Arrival_Airport:  searchDetails.ArrivalAirport,
		}
		return cf, nil
	}

	check, returnFlight := GetFlightFacilities(svc, returnPath, selectReq)
	if !check {
		return nil, errors.New("unable to fetch direct and return flight details")
	}
	cf := &pb.CompleteFlightDetails{
		DirectFlight:     directFlight,
		ReturnFlight:     returnFlight,
		NumberOfAdults:   selectReq.Adults,
		NumberOfChildren: selectReq.Children,
		CabinClass:       cabinClass,
		DepartureAirport: searchDetails.DepartureAirport,
		Arrival_Airport:  searchDetails.ArrivalAirport,
	}
	return cf, nil
}

func GetFlightFacilities(svc *AdminAirlineServiceStruct, path *pb.SearchFlightDetails, addDetails *pb.SelectFlightAdmin) (bool, *pb.FlightFacilities) {
	if len(path.FlightSegment) == 0 {
		log.Println("flight path is zero")
		return false, nil
	}

	flightNumber := path.FlightSegment[0].FlightNumber
	flight, err := svc.repo.FindFlightByFlightNumber(flightNumber)
	if err != nil {
		log.Println("unable to get flight details GetFlightFacilities() - kafka_services")
		return false, nil
	}

	cancellation, err := getCancellationDetails(svc, *flight)
	baggage, err := getBaggageDetails(svc, *flight)
	seatCheckD, _ := seatAvailabilityCheck(svc, path.FlightSegment, addDetails)
	fare := path.FlightSegment[0].FlightFare
	//seatCheckR, ReturnSeats := seatAvailabilityCheck(svc, path.Flights, addDetails)
	if !seatCheckD {
		return false, nil
	}

	cancel := &pb.Cancellation{
		CancellationPercentage:     float32(cancellation.CancellationPercentage),
		CancellationDeadlineBefore: int32(cancellation.CancellationDeadlineBefore),
		Refundable:                 cancellation.Refundable,
	}

	bag := &pb.Baggage{
		CabinAllowedHeight:  float32(baggage.CabinAllowedHeight),
		CabinAllowedWeight:  float32(baggage.CabinAllowedWeight),
		CabinAllowedBreadth: float32(baggage.CabinAllowedBreadth),
		CabinAllowedLength:  float32(baggage.CabinAllowedLength),
		HandAllowedHeight:   float32(baggage.HandAllowedHeight),
		HandAllowedWeight:   float32(baggage.HandAllowedWeight),
		HandAllowedBreadth:  float32(baggage.HandAllowedBreadth),
		HandAllowedLength:   float32(baggage.HandAllowedLength),
		FeeExtraPerKgCabin:  float32(baggage.FeeExtraPerKGHand),
		FeeExtraPerKgHand:   float32(baggage.FeeExtraPerKGCabin),
		Restrictions:        baggage.Restrictions,
	}
	return true, &pb.FlightFacilities{
		Cancellation: cancel,
		Baggage:      bag,
		FlightPath:   path,
		Fare:         fare,
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

func seatAvailabilityCheck(svc *AdminAirlineServiceStruct, flights []*pb.FlightDetails, addDetails *pb.SelectFlightAdmin) (bool, []dom.TemporaryData) {
	var tempData []dom.TemporaryData
	for _, f := range flights {
		chartID := f.FlightChartId
		seat, _ := svc.repo.FindSeatsByChartID(uint(chartID))
		if addDetails.Economy {
			if (int(addDetails.Adults) + int(addDetails.Children)) > seat.EconomySeatTotal-seat.EconomySeatBooked {
				log.Printf("economy seats in flight %v overbooked", f.FlightNumber)
				return false, nil
			}
		} else {
			log.Printf("business seats in flight %v overbooked", f.FlightNumber)
			if (int(addDetails.Adults) + int(addDetails.Children)) > seat.BusinessSeatTotal-seat.BusinessSeatBooked {
				return false, nil
			}
		}
		tempData = append(tempData, dom.TemporaryData{
			ChartId: uint(chartID),
			Seats:   seat,
		})
	}
	return true, tempData
}
