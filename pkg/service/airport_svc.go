package service

import (
	"context"
	"errors"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
)

func (svc *AdminAirlineServiceStruct) CreateAirport(p *pb.Airport) (*dom.Airport, error) {
	airport, err := svc.repo.FindAirportByAirportCode(p.AirportCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found  of model %v", p.AirportCode)
		} else {
			log.Printf("Airport not created of code %v, err: %v", p.AirportCode, err.Error())
			return airport, err
		}
	}

	airport, err = svc.repo.CreateAirport(p)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("Flight Type not created of model %v, err: %v",
				p.AirportCode, err.Error())
			return nil, err
		} else {
			return nil, err
		}
	}
	return airport, nil
}

func (svc *AdminAirlineServiceStruct) GetAirport(ctx context.Context, p *pb.AirportRequest) (*pb.AirportResponse, error) {
	resp, err := svc.repo.FindAirportByAirportCode(p.AirportCode)
	if err != nil {
		return nil, err
	}
	return &pb.AirportResponse{
		Airport: &pb.Airport{
			AirportCode:  resp.AirportCode,
			AirportName:  resp.AirportName,
			City:         resp.City,
			Country:      resp.Country,
			Region:       resp.Region,
			IcaoCode:     resp.ICAOCode,
			IataFcsCode:  resp.IATAFCSCode,
			Website:      resp.Website,
			ContactEmail: resp.ContactEmail,
			ContactPhone: resp.ContactPhone,
		},
	}, err
}
func (svc *AdminAirlineServiceStruct) GetAirports(ctx context.Context, p *pb.EmptyRequest) (*pb.AirportsResponse, error) {
	resp, err := svc.repo.FindAllAirports()
	if err != nil {
		return nil, err
	}
	return ConvertToResponseAirport(resp), err
}

func (svc *AdminAirlineServiceStruct) DeleteAirport(ctx context.Context, p *pb.AirportRequest) error {
	err := svc.repo.DeleteAirportByCode(p.AirportCode)
	if err != nil {
		return err
	}
	return nil
}

func ConvertToResponseAirport(data []dom.Airport) *pb.AirportsResponse {
	var result []*pb.Airport
	for _, resp := range data {
		result = append(result, &pb.Airport{
			AirportCode:  resp.AirportCode,
			AirportName:  resp.AirportName,
			City:         resp.City,
			Country:      resp.Country,
			Region:       resp.Region,
			IcaoCode:     resp.ICAOCode,
			IataFcsCode:  resp.IATAFCSCode,
			Website:      resp.Website,
			ContactEmail: resp.ContactEmail,
			ContactPhone: resp.ContactPhone,
		})
	}
	return &pb.AirportsResponse{
		Airports: result,
	}
}
