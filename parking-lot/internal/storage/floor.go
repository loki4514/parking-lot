package storage

type FloorIds string

const (
	F1 FloorIds = "F1"
	F2 FloorIds = "F2"
)

type Floor struct {
	FloorId     string
	FloorNumber int
	Spots       []Spot
}

// First floor: 50 bike, 50 bicycle, 25 car, 25 van
func FirstFloor() Floor {
	spots := GenerateSpots(50, 50, 25, 25)
	return Floor{
		FloorId:     string(F1),
		FloorNumber: 1,
		Spots:       spots,
	}
}

// Second floor: maybe different distribution
// Example: 40 bike, 40 bicycle, 20 car, 20 van
func SecondFloor() Floor {
	spots := GenerateSpots(40, 40, 20, 20)
	return Floor{
		FloorId:     string(F2),
		FloorNumber: 2,
		Spots:       spots,
	}
}
