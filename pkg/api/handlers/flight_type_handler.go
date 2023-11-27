package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/utils"
	"google.golang.org/grpc/metadata"
)

//* METHODS BELOW THIS IS TO HANDLE FLIGHT TYPES

func (handler *AdminAirlineHandler) RegisterFlightType(ctx context.Context, p *pb.FlightTypeRequest) (*pb.FlightTypeResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting")
		return nil, errors.New("Deadline Passed, returning")
	}

	response, err := handler.svc.CreateFlightType(p)
	if err != nil {
		log.Printf("Unable to create, err: %v", err.Error())
		return nil, err
	}
	flightTypeResponse := utils.ConvertFlightModelToResponse(response)
	return flightTypeResponse, nil
}

func (handler *AdminAirlineHandler) FetchAllFlightTypes(ctx context.Context, p *pb.EmptyRequest) (*pb.FlightTypesResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting")
		return nil, errors.New("Deadline Passed, returning")
	}

	response, err := handler.svc.GetAllFlightTypes()
	if err != nil {
		log.Println("Unable to fetch flight types some error in service")
		return nil, errors.New("Unable to fetch flight types")
	}
	flightTypesResponse := utils.ConvertFlightModelsToResponse(response)

	return flightTypesResponse, nil
}

func (handler *AdminAirlineHandler) FetchFlightType(ctx context.Context, p *pb.IDRequest) (*pb.FlightTypeResponse, error) {
	errChan := make(chan error)
	res := make(chan *pb.FlightTypeResponse)

	// go func() {
	// default: {
	go func() {
		fmt.Println("raed")

		//* testing using sleep
		// time.Sleep(time.Second * 5)

		md, check := metadata.FromIncomingContext(ctx)
		if !check {
			log.Println("unable to get flight_type_id from metada")
			// return nil, errors.New("unable to get flight_type_id from metada")
			res <- nil
			errChan <- errors.New("unable to get flight_type_id from metada")
		}

		flightId := md.Get("id")
		if len(flightId) == 0 {
			log.Println("flight id not received")
			// return nil, errors.New("flight id not received")
			res <- nil
			errChan <- errors.New("unable to get flight_type_id from metada")
		}

		time.Sleep(time.Second * 5)
		fmt.Println("majeed")

		id, _ := strconv.Atoi(flightId[0])
		response, err := handler.svc.GetFlightType(int32(id))

		if err != nil {
			log.Printf("Unable to fetch flight types err: %v", err.Error())
			// return nil, fmt.Errorf("Unable to fetch flight  types err: %v", err.Error())
			res <- nil
			errChan <- errors.New("unable to get flight_type_id from metada")
		}

		time.Sleep(time.Second * 5)
		fmt.Println("abdul")
		flightTypeResponse := utils.ConvertFlightModelToResponse(response)

		errChan <- nil
		close(errChan)
		res <- flightTypeResponse
	}()
	// return flightTypeResponse, nil

	select {
	case <-ctx.Done():
		fmt.Println("terminating")
		return nil, errors.New("terminated context kaunilla")
	case e := <-errChan:
		return <-res, e
	}
}

// }
// }
// }

// return flightTypeResponse, nil
// res <- flightTypeResponse
// errChan <- nil
// }()

// for {
// case <-errChan:
// 	return <-res, <-errChan
// }
// }

func (handler *AdminAirlineHandler) UpdateFlightType(ctx context.Context, p *pb.FlightTypeRequest) (*pb.FlightTypeResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting")
		return nil, errors.New("Deadline Passed, returning")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("flight id not received")
		return nil, errors.New("flight id not received")
	}

	flightId := md.Get("flight_type_id")
	if len(flightId) == 0 {
		log.Println("flight id not received")
		return nil, errors.New("flight id not received")
	}

	flight_id, err := strconv.Atoi(flightId[0])
	if err != nil {
		log.Printf("No id found or error converting string")
		return nil, errors.New("No id found or error converting string")
	}

	response, err := handler.svc.UpdateFlightType(p, flight_id)
	if err != nil {
		log.Printf("Unable to fetch flight types, err: %v", err.Error())
		return nil, fmt.Errorf("Unable to fetch flight types, err: %v", err.Error())
	}
	flightTypeResponse := utils.ConvertFlightModelToResponse(response)
	return flightTypeResponse, nil
}

func (handler *AdminAirlineHandler) DeleteFlightType(ctx context.Context, p *pb.IDRequest) (*pb.FlightTypeResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting")
		return nil, errors.New("Deadline Passed, returning")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("flight id not received")
		return nil, errors.New("flight id not received")
	}

	flightId := md.Get("flight_type_id")
	if len(flightId) == 0 {
		log.Println("flight id not received")
		return nil, errors.New("flight id not received")
	}

	flight_id, err := strconv.Atoi(flightId[0])
	if err != nil {
		log.Printf("No id found or error converting string")
		return nil, errors.New("No id found or error converting string")
	}

	response, err := handler.svc.DeleteFlightType(flight_id)
	if err != nil {
		log.Printf("Unable to fetch flight types, err: %v", err.Error())
		return nil, fmt.Errorf("Unable to fetch flight types, err: %v", err.Error())
	}
	flightTypeResponse := utils.ConvertFlightModelToResponse(response)
	return flightTypeResponse, nil
}
