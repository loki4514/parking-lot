package storage

import "time"

type VehicleType string

const (
	Car     VehicleType = "Car"
	Bike    VehicleType = "Bike"
	Bicycle VehicleType = "Bicycle"
	Van     VehicleType = "Van"
)

type Vehicle struct {
	VehicleType        VehicleType
	VehicleNumberPlate string
	SpotId             string
	EntryTime          time.Time
	ExitTime           time.Time
	IsPaid             bool
}
