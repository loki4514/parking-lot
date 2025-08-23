package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/loki4514/parking-lot/internal/storage"
)

// CreateEntryTicket creates a ticket for a vehicle and marks the spot as occupied
func CreateEntryTicket(vehicle storage.VehicleType, parkingLot *storage.ParkingLot) (*storage.Ticket, error) {
	// Find an available spot for the vehicle
	floorId, spotId, err := getSpot(vehicle, parkingLot)
	if err != nil {
		return nil, err
	}

	// Generate a unique Vehicle ID
	vehId := uuid.New().String()

	// Create the ticket
	ticket := storage.CreateTicket(spotId, floorId, vehId, vehicle)

	// Mark the spot as occupied
	for fi := range parkingLot.Floors {
		if parkingLot.Floors[fi].FloorId == floorId {
			for si := range parkingLot.Floors[fi].Spots {
				if parkingLot.Floors[fi].Spots[si].SpotId == spotId {
					parkingLot.Floors[fi].Spots[si].Availability = false
					break
				}
			}
			break
		}
	}

	fmt.Printf("Ticket created successfully: %+v\n", ticket)

	return ticket, nil
}
