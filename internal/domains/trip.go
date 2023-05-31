package domains

import "time"

type Trip struct {
	Id               int       `json:"id"`
	DriverId         int       `json:"driverId"`
	CarId            int       `json:"carId"`
	CurrentOccupancy int       `json:"currentOccupancy"`
	DepartureTime    time.Time `json:"departureTime"`
	DepartureCity    string    `json:"departureCity"`
	DestinationCity  string    `json:"destinationCity"`
	Price            float32   `json:"price"`
}

type Booking struct {
	Id            int `json:"id"`
	TripId        int `json:"tripId"`
	UserId        int `json:"userId"`
	SpotsOccupied int `json:"spotsOccupied"`
}
