package services

import (
	"errors"
	"time"

	"github.com/loki4514/parking-lot/internal/storage"
)

func FindTicket(ticketId string) (string, string, storage.VehicleType, time.Time, error) {
	for _, ticket := range storage.Tickets {
		if ticket.TicketID == ticketId {
			return ticket.FloorID, ticket.SpotID, ticket.VehicleType, ticket.EntryTime, nil
		}
	}
	var vt storage.VehicleType
	return "", "", vt, time.Now(), errors.New("invalid ticket id, please entry the valid tickets id")

}

func UpdateTicket(ticketId string, exitTime time.Time, fee int) (*storage.Ticket, error) {
	// Check if the ticket exists
	ticket, exists := storage.Tickets[ticketId]
	if !exists {
		return nil, errors.New("ticket not found")
	}

	// Update fields
	ticket.ExitTime = exitTime
	ticket.Fee = float64(fee)
	ticket.PaymentStatus = true // Assuming payment is done

	// Return updated ticket
	return ticket, nil
}

func UpdatingSpot(floorId string, spotId string, vehicle storage.VehicleType, floors *storage.ParkingLot) error {
	// Step 1: Find the floor
	var floorIndex = -1
	for i, floor := range floors.Floors {
		if floor.FloorId == floorId {
			floorIndex = i
			break
		}
	}
	if floorIndex == -1 {
		return errors.New("floor not found")
	}

	// Step 2: Find the spot
	var spotFound = false
	for i := range floors.Floors[floorIndex].Spots {
		spot := &floors.Floors[floorIndex].Spots[i]
		if spot.SpotId == spotId {
			// Step 3: Check vehicle type
			if spot.SpotVehicleType != vehicle {
				return errors.New("vehicle type mismatch")
			}
			// Step 4: Mark spot as available
			spot.Availability = true
			spotFound = true
			break
		}
	}

	// Step 5: Spot not found
	if !spotFound {
		return errors.New("spot not found")
	}

	return nil
}
