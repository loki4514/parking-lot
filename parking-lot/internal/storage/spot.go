package storage

import (
	"strconv"
)

type Spot struct {
	SpotId          string
	SpotVehicleType VehicleType
	Availability    bool
}

// global slice of all spots
func GenerateBikeSpots(count int) []Spot {
	var spots []Spot
	for i := 1; i <= count; i++ {
		spot := Spot{
			SpotId:          "BIKE-" + strconv.Itoa(i),
			SpotVehicleType: Bike,
			Availability:    true,
		}
		spots = append(spots, spot)
	}
	return spots
}

func GenerateBicycleSpots(count int) []Spot {
	var spots []Spot
	for i := 1; i <= count; i++ {
		spot := Spot{
			SpotId:          "BICYCLE-" + strconv.Itoa(i),
			SpotVehicleType: Bicycle,
			Availability:    true,
		}
		spots = append(spots, spot)
	}
	return spots
}

func GenerateCarSpots(count int) []Spot {
	var spots []Spot
	for i := 1; i <= count; i++ {
		spot := Spot{
			SpotId:          "CAR-" + strconv.Itoa(i),
			SpotVehicleType: Car,
			Availability:    true,
		}
		spots = append(spots, spot)
	}
	return spots
}

func GenerateVanSpots(count int) []Spot {
	var spots []Spot
	for i := 1; i <= count; i++ {
		spot := Spot{
			SpotId:          "VAN-" + strconv.Itoa(i),
			SpotVehicleType: Van,
			Availability:    true,
		}
		spots = append(spots, spot)
	}
	return spots
}

// One function to generate all types
func GenerateSpots(bikeCount, cycleCount, carCount, vanCount int) []Spot {
	var allSpots []Spot
	allSpots = append(allSpots, GenerateBikeSpots(bikeCount)...)
	allSpots = append(allSpots, GenerateBicycleSpots(cycleCount)...)
	allSpots = append(allSpots, GenerateCarSpots(carCount)...)
	allSpots = append(allSpots, GenerateVanSpots(vanCount)...)
	return allSpots
}
