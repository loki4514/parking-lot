# 🅿️ Parking Lot Management System

A comprehensive parking lot management system built in Go that implements various design patterns and provides a CLI interface for vehicle parking operations.

## 📋 Table of Contents
- [Features](#features)
- [System Architecture](#system-architecture)
- [Installation](#installation)
- [Usage](#usage)
- [Design Patterns](#design-patterns)
- [Project Structure](#project-structure)
- [API Reference](#api-reference)
- [Configuration](#configuration)
- [Contributing](#contributing)

## ✨ Features

- **Multi-Vehicle Support**: Cars, Bikes, Bicycles, and Vans
- **Multi-Floor Parking**: 2-floor parking lot with configurable spots
- **Real-time Availability**: Dynamic spot allocation and tracking
- **Fee Calculation**: Strategy pattern-based pricing for different vehicle types
- **Ticket Management**: UUID-based ticket generation and validation
- **Payment Processing**: Cash and Card payment options
- **CLI Interface**: User-friendly command-line interface

## 🏗️ System Architecture

### Vehicle Types & Pricing
| Vehicle Type | Hourly Rate | Spots (Floor 1) | Spots (Floor 2) |
|-------------|-------------|-----------------|-----------------|
| Bicycle     | ₹20/hour    | 50              | 40              |
| Bike        | ₹50/hour    | 50              | 40              |
| Car         | ₹100/hour   | 25              | 20              |
| Van         | ₹150/hour   | 25              | 20              |

### Parking Lot Layout
```
📊 PARKING LOT CAPACITY
┌─────────────┬──────────┬──────────┬───────┐
│ Vehicle     │ Floor 1  │ Floor 2  │ Total │
├─────────────┼──────────┼──────────┼───────┤
│ Bicycles    │    50    │    40    │   90  │
│ Bikes       │    50    │    40    │   90  │
│ Cars        │    25    │    20    │   45  │
│ Vans        │    25    │    20    │   45  │
├─────────────┼──────────┼──────────┼───────┤
│ **Total**   │  **150** │  **120** │ **270** │
└─────────────┴──────────┴──────────┴───────┘
```

## 🚀 Installation

### Prerequisites
- Go 1.19 or higher
- Git

### Setup
```bash
# Clone the repository
git clone https://github.com/loki4514/parking-lot.git
cd parking-lot

# Initialize Go modules
go mod init github.com/loki4514/parking-lot
go mod tidy

# Install dependencies
go get github.com/google/uuid

# Run the application
go run main.go
```

## 💻 Usage

### Starting the System
```bash
go run main.go
```

### Available Commands
```
Available Commands:
ENTER → Create Entry Ticket
EXIT  → Exit Parking Lot  
Q     → Quit
```

### Example Workflow

#### 1. Vehicle Entry
```
Enter your command: enter
Available Vehicle Types: Car, Bike, Bicycle, Van
Enter your vehicle type: car
✅ Entry ticket created successfully! Ticket ID: 550e8400-e29b-41d4-a716-446655440000
```

#### 2. Vehicle Exit
```
Enter your command: exit
Enter your Ticket ID: 550e8400-e29b-41d4-a716-446655440000
💰 Parking fee: ₹150
Payment method (CASH / CARD): CASH
✅ Payment successful via CASH.
✅ Exit successful! Thank you for visiting.
```

## 🎨 Design Patterns

### 1. Strategy Pattern
Used for fee calculation based on vehicle type:
```go
type FeeCalculation interface {
    CalculateFee(entryTime, exitTime time.Time) int
}

// Implementations: BicycleFeeCalculation, BikeFeeCalculation, etc.
```

### 2. Factory Pattern
Spot generation for different vehicle types:
```go
func GenerateSpots(bikeCount, cycleCount, carCount, vanCount int) []Spot
```

### 3. Repository Pattern
Ticket storage and retrieval operations:
```go
var Tickets = map[string]*Ticket{}
func CreateTicket(spotId, floorId, vehicleID string, vehicleType VehicleType) *Ticket
```

## 📁 Project Structure

```
parking-lot/
├── main.go                     # Main application entry point
├── internal/
│   ├── storage/
│   │   ├── floor.go           # Floor structure and initialization
│   │   ├── parking_lot.go     # Main parking lot structure
│   │   ├── spot.go           # Parking spot management
│   │   ├── ticket.go         # Ticket structure and operations
│   │   └── vehicle.go        # Vehicle types and definitions
│   └── services/
│       ├── entry_ticket.go   # Entry ticket creation service
│       ├── fee_calculation.go # Strategy pattern for fee calculation
│       ├── spot_finder.go    # Available spot finding logic
│       └── ticket_service.go # Ticket management operations
├── go.mod
├── go.sum
└── README.md
```

## 📚 API Reference

### Core Data Structures

#### VehicleType
```go
type VehicleType string
const (
    Car     VehicleType = "Car"
    Bike    VehicleType = "Bike" 
    Bicycle VehicleType = "Bicycle"
    Van     VehicleType = "Van"
    Unknown VehicleType = "Unknown"
)
```

#### Ticket
```go
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
```

#### Spot
```go
type Spot struct {
    SpotId          string
    SpotVehicleType VehicleType
    Availability    bool
}
```

### Key Functions

#### Entry Operations
- `CreateEntryTicket(vehicle VehicleType, parkingLot *ParkingLot) (*Ticket, error)`
- `getSpot(vehicle VehicleType, parkingLot *ParkingLot) (string, string, error)`

#### Exit Operations
- `FindTicket(ticketId string) (string, string, VehicleType, time.Time, error)`
- `UpdateTicket(ticketId string, exitTime time.Time, fee int) (*Ticket, error)`
- `UpdatingSpot(floorId, spotId string, vehicle VehicleType, parkingLot *ParkingLot) error`

#### Fee Calculation
- `StrategyPattern(vehicleType VehicleType, entryTime, exitTime time.Time) int`

## ⚙️ Configuration

### Customizing Spot Distribution
Modify the floor generation functions in `storage/floor.go`:

```go
// First floor: 50 bike, 50 bicycle, 25 car, 25 van
func FirstFloor() Floor {
    spots := GenerateSpots(50, 50, 25, 25)
    // ... rest of the implementation
}
```

### Customizing Fee Structure
Modify the base rates in `services/fee_calculation.go`:

```go
// Bicycle: ₹20/hour
func (b *BicycleFeeCalculation) CalculateFee(entryTime, exitTime time.Time) int {
    return estimateFare(20, entryTime, exitTime)
}
```

## 🔧 Improvements & Future Enhancements

### Current Limitations
- ⚠️ **Thread Safety**: Not safe for concurrent access
- ⚠️ **Data Persistence**: Uses in-memory storage only
- ⚠️ **Error Recovery**: Limited error handling for edge cases

### Planned Features
- [ ] Database integration (PostgreSQL/MongoDB)
- [ ] REST API endpoints
- [ ] Web dashboard for parking lot management
- [ ] Real-time parking availability notifications
- [ ] Vehicle reservation system
- [ ] Integration with payment gateways
- [ ] Parking analytics and reporting
- [ ] Multi-tenant support for multiple parking lots

### Technical Improvements
- [ ] Add mutex locks for thread safety
- [ ] Implement proper logging with structured logs
- [ ] Add configuration file support (YAML/JSON)
- [ ] Unit and integration tests
- [ ] Docker containerization
- [ ] Health check endpoints
- [ ] Rate limiting for API calls

## 🧪 Testing

```bash
# Run tests (when implemented)
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/services
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Code Style Guidelines
- Follow standard Go formatting (`go fmt`)
- Use meaningful variable and function names
- Add comments for complex logic
- Maintain consistent error handling patterns

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

**loki4514**
- GitHub: [@loki4514](https://github.com/loki4514)

## 🙏 Acknowledgments

- Thanks to the Go community for excellent documentation
- Inspired by real-world parking management systems
- Design patterns reference from Gang of Four

---

⭐ **Star this repo if you found it helpful!**