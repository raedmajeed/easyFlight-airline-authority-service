package service

import (
	"encoding/json"
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

func (svc *AdminAirlineServiceStruct) CalculateDailyFare() {
	bookedSeats, err := svc.repo.FindAllBookedSeats()
	if err != nil {
		log.Println("error fetching booked seats for fair details")
	}

	for _, seat := range bookedSeats {
		id := seat.FlightChartNo
		calculateAndSavePriceDaily(svc, id, &seat)
	}
}

func calculateAndSavePriceDaily(svc *AdminAirlineServiceStruct, id int, seats *dom.BookedSeat) {
	flightChart, err := svc.repo.FindFlightChartById(id)
	response, err := svc.repo.FindScheduleByID(int(flightChart.ScheduleID))
	//seats, err := svc.repo.FindSeatsByChartID(flightID)
	if err != nil {
		log.Printf("unable to get schedule ID, in method  calculateAndSavePrice() - service, err: %v", err.Error())
		return
	}

	departureDate := response.DepartureDateTime
	departureAirport := response.DepartureAirport
	arrivalAirport := response.ArrivalAirport
	todayDate := time.Now()
	remainingDays := departureDate.Sub(todayDate)
	days := int(remainingDays.Hours() / 24)
	onlyDate := todayDate.Format("2006-01-02")
	businessSurgeFactor, _ := strconv.ParseFloat(svc.cfg.BUSINESSSURGE, 64)

	if departureAirport == "" {
		return
	}
	depResponse, err := svc.repo.FindAirportByAirportCode(departureAirport)
	if err != nil {
		log.Printf("unable to get departure airport, in method  calculateAndSavePrice() - service, err: %v", err.Error())
		return
	}
	ArrResponse, err := svc.repo.FindAirportByAirportCode(arrivalAirport)
	if err != nil {
		log.Printf("unable to get arrival airport, in method  calculateAndSavePrice() - service, err: %v", err.Error())
		return
	}
	// take the schedule and find how many days left for departure

	DaysLeftPercentage := CalculateCustomPercentage(days)
	// find schedule here and calculate the distance
	distance := DistanceCalculator(depResponse.Latitude, depResponse.Longitude, ArrResponse.Latitude, ArrResponse.Longitude)
	// once I get the distance fetch today's petrol price
	fuelPrice, err := FuelPricedDaily() //DONE
	// check if today's date has any holiday
	holidayPercentage := dom.Holidays(onlyDate)
	// check % of the days
	weekDayPercentage := dom.DaysOFWeek()
	// check how many seats booked, if % > 50 adjust price accordingly
	eFare, bFare := SeatsBookedPercentage(seats)
	// finally once I get all the values add the price
	EconomyFare := (fuelPrice * distance) / 2
	EconomyFare = EconomyFare + ((EconomyFare * DaysLeftPercentage) / 100)
	EconomyFare = EconomyFare + ((EconomyFare * holidayPercentage) / 100)
	EconomyFare = EconomyFare + ((EconomyFare * weekDayPercentage) / 100)
	EconomyFare = EconomyFare + ((EconomyFare * eFare) / 100)

	BusinessFare := (fuelPrice * distance) / 2
	BusinessFare = BusinessFare + ((BusinessFare * DaysLeftPercentage) / 100)
	BusinessFare = BusinessFare + ((BusinessFare * holidayPercentage) / 100)
	BusinessFare = BusinessFare + ((BusinessFare * weekDayPercentage) / 100)
	BusinessFare = BusinessFare + ((BusinessFare * bFare) / 100)
	BusinessFare = BusinessFare * businessSurgeFactor

	flightChart.EconomyFare = EconomyFare / 10
	flightChart.BusinessFare = BusinessFare / 10
	err = svc.repo.UpdateFlightChart(flightChart)
	if err != nil {
		log.Println("unable to update flight fare")
		return
	}
}

func FuelPricedDaily() (float64, error) {
	res, err := http.Get("https://mfapps.indiatimes.com/ET_Calculators/oilpricebymetro.htm")
	if err != nil {
		log.Println("unable to fetch today's petrol price")
		return 100, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("unable to read response, in method FuelPricedDaily() - service, err: %v", err.Error())
		return 100, err
	}

	var prices dom.PetrolPrice
	err = json.Unmarshal(body, &prices)
	if err != nil {
		log.Printf("unable to marshal json, in method  FuelPricedDaily() - service, err: %v", err.Error())
		return 100, err
	}

	todayPetrolPrice := prices.Prices[0].PetrolPrice
	price, err := strconv.ParseFloat(todayPetrolPrice, 64)
	if err != nil {
		log.Println(err)
		return 100, err
	}
	return price, nil
}

func SeatsBookedPercentage(seats *dom.BookedSeat) (float64, float64) {
	eFare := 0
	bFare := 0
	economySeatsBooked := seats.EconomySeatBooked
	economySeatsTotal := seats.EconomySeatTotal
	businessSeatsBooked := seats.BusinessSeatBooked
	businessSeatsTotal := seats.BusinessSeatTotal

	ePercentage := (economySeatsBooked / economySeatsTotal) * 100
	bPercentage := (businessSeatsBooked / businessSeatsTotal) * 100

	if ePercentage > 50 {
		eFare = 7
	} else if ePercentage > 70 {
		eFare = 10
	} else if ePercentage > 90 {
		eFare = 14
	}

	if bPercentage > 50 {
		bFare = 7
	} else if bPercentage > 70 {
		bFare = 10
	} else if bPercentage > 90 {
		bFare = 14
	}

	return float64(eFare), float64(bFare)
}

func CalculateCustomPercentage(days int) float64 {
	maxPercentage := 15.0
	percentage := (float64(15-days) / 14) * maxPercentage

	if percentage > maxPercentage {
		percentage = maxPercentage
	}
	return percentage
}

func DistanceCalculator(lat1, lon1, lat2, lon2 float64) float64 {
	// Convert latitude and longitude from degrees to radians
	const EarthRadius = 6371
	lat1Rad := lat1 * (math.Pi / 180)
	lon1Rad := lon1 * (math.Pi / 180)
	lat2Rad := lat2 * (math.Pi / 180)
	lon2Rad := lon2 * (math.Pi / 180)

	// Haversine formula
	dlon := lon2Rad - lon1Rad
	dlat := lat2Rad - lat1Rad
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Calculate the distance
	distance := EarthRadius * c
	return distance
}
