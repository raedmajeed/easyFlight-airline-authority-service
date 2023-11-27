package model

import "time"

type Holiday struct {
	Name    string
	Date    time.Time
	Percent float64
}

type DayOfWeekPercent struct {
	DayOfWeek time.Weekday
	Percent   float64
}

type FuelPrice struct {
	CityState    string `json:"cityState"`
	PetrolPrice  string `json:"petrolPrice"`
	CreatedDate  string `json:"createdDate"`
	PriceDate    string `json:"priceDate"`
	PetrolChange string `json:"petrolChange"`
	DieselPrice  string `json:"dieselPrice"`
	DieselChange string `json:"dieselChange"`
	Origin       string `json:"origin"`
	ID           int    `json:"ID"`
	IsActive     int    `json:"isActive"`
	Type         string `json:"type"`
	Seoname      string `json:"seoname"`
}

type PetrolPrice struct {
	Prices []FuelPrice `json:"results"`
}

func Holidays(date string) float64 {
	holiday := map[string]Holiday{
		"2023-11-11": {Name: "Diwali", Date: time.Date(2023, 11, 11, 0, 0, 0, 0, time.UTC), Percent: 11.5},
		"2023-12-25": {Name: "Christmas", Date: time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC), Percent: 12.0},
		"2024-01-01": {Name: "New Year's Day", Date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Percent: 10.75},
		"2024-04-07": {Name: "Easter Sunday", Date: time.Date(2024, 4, 7, 0, 0, 0, 0, time.UTC), Percent: 8.25},
		"2024-07-04": {Name: "Independence Day (USA)", Date: time.Date(2024, 7, 4, 0, 0, 0, 0, time.UTC), Percent: 5.25},
		"2024-10-31": {Name: "Halloween", Date: time.Date(2024, 10, 31, 0, 0, 0, 0, time.UTC), Percent: 14.15},
		"2024-11-11": {Name: "Diwali", Date: time.Date(2024, 11, 11, 0, 0, 0, 0, time.UTC), Percent: 21.5},
	}
	v, ok := holiday[date]
	if !ok {
		return 0
	}
	return v.Percent
}

func DaysOFWeek() float64 {
	weekDays := map[time.Weekday]DayOfWeekPercent{
		time.Sunday:    {DayOfWeek: time.Sunday, Percent: 4.25},
		time.Saturday:  {DayOfWeek: time.Saturday, Percent: 2.15},
		time.Friday:    {DayOfWeek: time.Friday, Percent: 1.9},
		time.Monday:    {DayOfWeek: time.Monday, Percent: 1.0},
		time.Tuesday:   {DayOfWeek: time.Tuesday, Percent: 1.0},
		time.Wednesday: {DayOfWeek: time.Wednesday, Percent: 1.0},
		time.Thursday:  {DayOfWeek: time.Thursday, Percent: 1.0},
	}

	v, ok := weekDays[time.Now().Weekday()]
	if !ok {
		return 0
	}
	return v.Percent
}
