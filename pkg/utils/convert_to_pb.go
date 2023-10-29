package utils

import (
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

func ConvertFlightModelToResponse(model *dom.FlightTypeModel) *pb.FlightTypeResponse {
	flightTypeRequest := &pb.FlightTypeRequest{
		Id:                  int32(model.ID),
		Type:                convertType(model.Type),
		FlightModel:         model.FlightModel,
		Description:         model.Description,
		ManufacturerName:    model.ManufacturerName,
		ManufacturerCountry: model.ManufacturerCountry,
		MaxDistance:         model.MaxDistance,
		CruiseSpeed:         model.CruiseSpeed,
	}

	flightTypeResponse := &pb.FlightTypeResponse{
		FlightType: flightTypeRequest,
	}

	return flightTypeResponse
}

func ConvertFlightModelsToResponse(models []dom.FlightTypeModel) *pb.FlightTypesResponse {
	fr := make([]*pb.FlightTypeRequest, len(models))
	for i := range models {
		pb := &pb.FlightTypeRequest{
			Id:                  int32(models[i].ID),
			Type:                convertType(models[i].Type),
			FlightModel:         models[i].FlightModel,
			Description:         models[i].Description,
			ManufacturerName:    models[i].ManufacturerName,
			ManufacturerCountry: models[i].ManufacturerCountry,
			MaxDistance:         models[i].MaxDistance,
			CruiseSpeed:         models[i].CruiseSpeed,
		}
		fr[i] = pb
	}
	return &pb.FlightTypesResponse{
		FlightTypes: fr,
	}
}

func convertType(typeStr string) pb.FlightTypeEnum {
	switch typeStr {
	case "Commercial":
		return pb.FlightTypeEnum_COMMERCIAL
	case "Military":
		return pb.FlightTypeEnum_MILITARY
	case "Cargo":
		return pb.FlightTypeEnum_CARGO
	default:
		return pb.FlightTypeEnum_COMMERCIAL
	}
}

func ConvertAirlineToResponse(model *dom.Airline) *pb.AirlineResponse {
	airlineRequest := &pb.AirlineRequest{
		Id:                   int32(model.ID),
		AirlineName:          model.AirlineName,
		CompanyAddress:       model.CompanyAddress,
		PhoneNumber:          model.PhoneNumber,
		Email:                model.Email,
		AirlineCode:          model.AirlineCode,
		AirlineLogoLink:      model.AirlineLogoLink,
		SupportDocumentsLink: model.SupportDocumentLink,
	}

	airlineResponse := &pb.AirlineResponse{
		Airline: airlineRequest,
	}

	return airlineResponse
}

func ConvertAirlineSeatsToResponse(model *dom.AirlineSeat) *pb.AirlineSeatResponse {
	airlineRequest := &pb.AirlineSeatRequest{
		AirlineId:           int32(model.AirlineId),
		EconomySeatNo:       int32(model.EconomySeatNumber),
		BuisinesSeatNo:      int32(model.BuisinesSeatNumber),
		EconomySeatsPerRow:  int32(model.EconomySeatsPerRow),
		BuisinesSeatsPerRow: int32(model.EconomySeatsPerRow),
	}

	airlineResponse := &pb.AirlineSeatResponse{
		AirlineSeat: airlineRequest,
	}

	return airlineResponse
}

func ConvertAirlineBaggagePolicyToResponse(model *dom.AirlineBaggage) *pb.AirlineBaggageResponse {
	baggageReq := &pb.AirlineBaggageRequest{
		AirlineId:           int32(model.AirlineId),
		Class:               convertTypeClass(model.FareClass),
		CabinAllowedWeight:  int32(model.CabinAllowedWeight),
		CabinAllowedLength:  int32(model.CabinAllowedLength),
		CabinAllowedBreadth: int32(model.CabinAllowedBreadth),
		CabinAllowedHeight:  int32(model.CabinAllowedHeight),
		HandAllowedWeight:   int32(model.HandAllowedWeight),
		HandAllowedLength:   int32(model.HandAllowedLength),
		HandAllowedBreadth:  int32(model.HandAllowedBreadth),
		HandAllowedHeight:   int32(model.HandAllowedHeight),
		FeeForExtraKgCabin:  int32(model.FeeExtraPerKGCabin),
		FeeForExtraKgHand:   int32(model.FeeExtraPerKGHand),
		Restrictions:        model.Restrictions,
	}
	return &pb.AirlineBaggageResponse{
		AirlineBaggage: baggageReq,
	}
}

func ConvertAirlineCancellationPolicyToResponse(model *dom.AirlineCancellation) *pb.AirlineCancellationResponse {
	cancellation := &pb.AirlineCancellationRequest{
		AirlineId:                       int32(model.AirlineId),
		Class:                           convertTypeClass(model.FareClass),
		CancellationDeadlineBeforeHours: uint32(model.CancellationDeadlineBefore),
		CancellationPercentage:          int32(model.CancellationPercentage),
		Refundable:                      model.Refundable,
	}

	return &pb.AirlineCancellationResponse{
		AirlineCancellation: cancellation,
	}
}

func convertTypeClass(typeStr int) pb.Class {
	switch typeStr {
	case 0:
		return pb.Class_ECONOMY
	case 1:
		return pb.Class_BUSINESS
	default:
		return pb.Class_ECONOMY
	}
}

func ConvertAirportToResponse(model *dom.Airport) *pb.AirportResponse {
	airportReq := &pb.Airport{
		AirportCode:  model.AirportCode,
		AirportName:  model.AirportName,
		City:         model.City,
		Country:      model.Country,
		Region:       model.Region,
		Latitude:     model.Latitude,
		Longitude:    model.Longitude,
		IataFcsCode:  model.IATAFCSCode,
		IcaoCode:     model.ICAOCode,
		Website:      model.Website,
		ContactEmail: model.ContactEmail,
		ContactPhone: model.ContactPhone,
	}
	return &pb.AirportResponse{
		Airport: airportReq,
	}
}

func ConvertSchedulesToResponse(model *dom.Schedule) *pb.ScheduleResponse {
	schedule := &pb.ScheduleRequest{
		DepartureAirport: model.DepartureAirport,
		ArrivalAirport:   model.ArrivalAirport,
		DepartureTime:    model.DepartureTime,
		ArrivalTime:      model.ArrivalTime,
	}

	return &pb.ScheduleResponse{
		Schedule: schedule,
	}
}

func ConvertLoginRequestToResponse(token string, p *pb.LoginRequest) *pb.LoginResponse {
	return &pb.LoginResponse{
		Email: p.Email,
		Token: token,
	}
}
