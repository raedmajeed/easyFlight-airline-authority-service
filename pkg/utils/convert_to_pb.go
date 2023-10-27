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
	// fmt.Println(fr)
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
		return pb.FlightTypeEnum_COMMERCIAL // Default value, change as needed
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
