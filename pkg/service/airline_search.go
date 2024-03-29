package service

import (
	"container/list"
	"fmt"
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	"log"
	"strconv"
	"time"
)

type FinalFlightList struct {
	List []*dom.FlightChart
}

// SearchFlight this function initiate the flight search
// checks if the flight is a direct flight or return
// and finds the connected and direct flight
func (svc *AdminAirlineServiceStruct) SearchFlight(search dom.SearchDetails) ([]dom.Path, []dom.Path, error) {
	var finalPaths []dom.Path
	var returnFinalPaths []dom.Path
	var err error

	if !search.ReturnFlight {
		finalPaths, err = OneWayFlightSearch(svc, search)
		return finalPaths, nil, err
	} else {
		finalPaths, err = OneWayFlightSearch(svc, search)
		temp := search.DepartureAirport
		search.DepartureAirport = search.ArrivalAirport
		search.ArrivalAirport = temp
		search.DepartureDate = search.ReturnDepartureDate
		returnFinalPaths, err = OneWayFlightSearch(svc, search)
	}
	log.Println("good byes", err)
	return finalPaths, returnFinalPaths, err
}

// OneWayFlightSearch this takes in the service, search details and finds
// the flight details connected and direct and sets into a
// path so that this can be used to fetch the details after
func OneWayFlightSearch(svc *AdminAirlineServiceStruct, searchInfo dom.SearchDetails) ([]dom.Path, error) {
	var finalFlightList []*dom.FlightChart
	var tempFlightList []*dom.FlightChart
	var tempFlightList2 []*dom.FlightChart

	flightLists, err := svc.repo.FindFlightsFromDep(searchInfo.DepartureAirport, searchInfo.DepartureDate)
	if err != nil {
		log.Println("error fetching flight path")
		return []dom.Path{}, fmt.Errorf("error fetching flight path OneWayFlightSearch(), err: %v", err.Error())
	}

	finalFlightList = addingToFinalFlightList(finalFlightList, flightLists)
	//// STOP - 1 search all arrival airports from previous departure airport
	for _, flight := range flightLists {
		chart, _ := svc.repo.FindScheduleByID(int(flight.ScheduleID))
		NewDepTime := chart.ArrivalDateTime.Add(time.Hour)
		NewDepAirport := chart.ArrivalAirport
		flightChart, _ := svc.repo.FindFlightScheduleByAirport(NewDepAirport, NewDepTime, int(chart.ID))
		finalFlightList = addingToFinalFlightList(finalFlightList, flightChart)
		tempFlightList = addingToFinalFlightList(tempFlightList, flightChart)
	}
	//// STOP - 2 search all arrival airports from previous departure airport
	for _, flight := range tempFlightList {
		chart, _ := svc.repo.FindScheduleByID(int(flight.ScheduleID))
		NewDepTime := chart.ArrivalDateTime.Add(time.Hour)
		NewDepAirport := chart.ArrivalAirport
		flightChart, _ := svc.repo.FindFlightScheduleByAirport(NewDepAirport, NewDepTime, int(chart.ID))
		finalFlightList = addingToFinalFlightList(finalFlightList, flightChart)
		tempFlightList2 = addingToFinalFlightList(tempFlightList2, flightChart)
	}

	//// Final search all arrival airport of last airport
	for _, flight := range tempFlightList2 {
		chart, _ := svc.repo.FindScheduleByID(int(flight.ScheduleID))
		NewDepTime := chart.ArrivalDateTime.Add(time.Hour)
		flightLists, _ = findAllFlightsFromDepAirport(svc, chart.ArrivalAirport, NewDepTime)
		finalFlightList = addingToFinalFlightList(finalFlightList, flightLists)
	}

	var finalFlightDetails []dom.FlightDetails
	for _, f := range finalFlightList {
		schedule, _ := svc.repo.FindScheduleByID(int(f.ScheduleID))
		flightFleet, _ := svc.repo.FindFlightFleetById(int(f.FlightID))
		airline, _ := svc.repo.FindAirlineById(int32(flightFleet.AirlineID))
		finalFlightDetails = append(finalFlightDetails, dom.FlightDetails{
			FlightChartID:     f.ID,
			FlightNumber:      f.FlightNumber,
			Airline:           airline.AirlineName,
			DepartureAirport:  schedule.DepartureAirport,
			ArrivalAirport:    schedule.ArrivalAirport,
			DepartureTime:     schedule.DepartureTime,
			ArrivalTime:       schedule.ArrivalTime,
			DepartureDate:     schedule.DepartureDate,
			ArrivalDate:       schedule.ArrivalDate,
			DepartureDateTime: schedule.DepartureDateTime,
			ArrivalDateTime:   schedule.ArrivalDateTime,
			EconomyFare:       f.EconomyFare,
			BusinessFare:      f.BusinessFare,
		})
	}
	//
	maxStops, _ := strconv.Atoi(searchInfo.MaxStops)
	//
	paths := FindAllPaths(searchInfo.DepartureAirport, searchInfo.ArrivalAirport, finalFlightDetails)
	finalResponsePaths := pathResponse(paths, maxStops)
	return finalResponsePaths, nil
}

func pathResponse(paths [][]dom.FlightDetails, maxStops int) []dom.Path {
	var finalPaths []dom.Path
	for i, path := range paths {
		if len(path)-1 > maxStops {
			continue
		}
		finalPaths = append(finalPaths, dom.Path{
			PathId:        i + 1,
			Flights:       path,
			NumberOfStops: len(path),
		})
	}
	return finalPaths
}

func FindAllPaths(departureAirport, arrivalAirport string, flights []dom.FlightDetails) [][]dom.FlightDetails {
	//* Create a graph to represent the flights.
	graph := make(map[string][]dom.FlightDetails)
	//* creating departure airport map here
	for _, flight := range flights {
		if _, ok := graph[flight.DepartureAirport]; !ok {
			graph[flight.DepartureAirport] = []dom.FlightDetails{}
		}
		graph[flight.DepartureAirport] = append(graph[flight.DepartureAirport], flight)
	}

	//Checking whether there is any flight from departure airport
	totalFlightsFromDep := graph[departureAirport]
	if len(totalFlightsFromDep) <= 0 {
		return [][]dom.FlightDetails{}
	}

	var finalPaths [][]dom.FlightDetails
	for _, flight := range totalFlightsFromDep {
		res := findPathsOfFlight(flight, arrivalAirport, graph)
		r := removeDuplicateLists(res)
		finalPaths = appendToFinalPath(finalPaths, r)
	}
	return finalPaths
}

func appendToFinalPath(f [][]dom.FlightDetails, r [][]dom.FlightDetails) [][]dom.FlightDetails {
	for _, i := range r {
		f = append(f, i)
	}
	return f
}

func removeDuplicateLists(lists [][]dom.FlightDetails) [][]dom.FlightDetails {
	uniqueLists := make(map[string]bool)
	var result [][]dom.FlightDetails

	for _, l := range lists {
		str := fmt.Sprintf("%v", l)
		if uniqueLists[str] {
			continue
		}
		uniqueLists[str] = true
		result = append(result, l)
	}

	return result
}

func findPathsOfFlight(start dom.FlightDetails, end string, graph map[string][]dom.FlightDetails) [][]dom.FlightDetails {
	var results [][]dom.FlightDetails
	queue := list.New()

	queue.PushBack([]dom.FlightDetails{start})
	visited := make(map[string]bool)

	for queue.Len() != 0 {
		pathElement := queue.Front()
		path := pathElement.Value.([]dom.FlightDetails)
		queue.Remove(pathElement)

		currentAirport := path[len(path)-1].ArrivalAirport

		if currentAirport == end {
			results = append(results, path)
			continue
		}

		for _, neighbor := range graph[currentAirport] {
			if !visited[neighbor.ArrivalAirport] {
				newPath := make([]dom.FlightDetails, len(path)+1)
				copy(newPath, path)
				newPath[len(newPath)-1] = neighbor

				if !path[len(path)-1].ArrivalDateTime.Add(time.Hour).Before(newPath[len(newPath)-1].DepartureDateTime) {
					continue
				}
				queue.PushBack(newPath)
			}
		}
		visited[currentAirport] = true
	}
	return results
}

func findAllFlightsFromDepAirport(svc *AdminAirlineServiceStruct, depAirport string, depDateTime time.Time) ([]*dom.FlightChart, error) {
	flightLists, err := svc.repo.FindFlightsFromAirport(depAirport, depDateTime)
	if err != nil {
		log.Println("unable to fetch flights")
		return nil, err
	}
	return flightLists, nil
}

func addingToFinalFlightList(lists []*dom.FlightChart, FinalFlightList []*dom.FlightChart) []*dom.FlightChart {
	for _, temp := range lists {
		FinalFlightList = append(FinalFlightList, temp)
	}
	return FinalFlightList
}
