package services

import (
	"errors"

	"github.com/loki4514/parking-lot/internal/storage"
)

func getSpot(vehicle storage.VehicleType, parkingLot *storage.ParkingLot) (string, string, error) {
	for _, floor := range parkingLot.Floors {
		for _, spot := range floor.Spots {
			if spot.SpotVehicleType == vehicle && spot.Availability {
				// Return immediately when a spot is found
				return floor.FloorId, spot.SpotId, nil
			}
		}
	}
	return "", "", errors.New("cannot find available spot for vehicle type or parking full")
}
