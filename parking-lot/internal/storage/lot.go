package storage

type Gate struct {
	GateID   string
	FloorID  string
	GateType string // "ENTRY" or "EXIT"
}

type ParkingLot struct {
	LotID      string
	Floors     []Floor
	EntryGates []Gate
	ExitGates  []Gate
}

// Create Parking Lot
func CreateParkingLot() ParkingLot {
	firstFloor := FirstFloor()
	secondFloor := SecondFloor()

	entryGates := []Gate{
		{GateID: "F1-Entry", FloorID: "F1", GateType: "ENTRY"},
		{GateID: "F2-Entry", FloorID: "F2", GateType: "ENTRY"},
	}

	exitGates := []Gate{
		{GateID: "F1-Exit", FloorID: "F1", GateType: "EXIT"},
		{GateID: "F2-Exit", FloorID: "F2", GateType: "EXIT"},
	}

	return ParkingLot{
		LotID:      "PARKINGLOT1",
		Floors:     []Floor{firstFloor, secondFloor},
		EntryGates: entryGates,
		ExitGates:  exitGates,
	}
}
