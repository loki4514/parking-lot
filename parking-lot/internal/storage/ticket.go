package storage

import (
	"time"

	"github.com/google/uuid"
)

// Ticket struct
type Ticket struct {
	TicketID      string
	FloorID       string
	SpotID        string
	VehicleID     string
	VehicleType   VehicleType
	EntryTime     time.Time
	ExitTime      time.Time
	Fee           float64
	PaymentStatus bool
}

// Global in-memory storage for tickets
var Tickets = map[string]*Ticket{}

// CreateTicket creates a new parking ticket and stores it in memory
func CreateTicket(spotId string, floorId string, vehicleID string, vehicleType VehicleType) *Ticket {
	ticketID := uuid.New().String()

	ticket := &Ticket{
		TicketID:    ticketID,
		FloorID:     floorId,
		SpotID:      spotId,
		VehicleID:   vehicleID,
		VehicleType: vehicleType,
		EntryTime:   time.Now(),
		// ExitTime, Fee, and PaymentStatus will be updated later when the vehicle leaves
	}

	// Store the ticket in the in-memory map
	Tickets[ticketID] = ticket

	return ticket
}
