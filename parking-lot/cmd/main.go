package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/loki4514/parking-lot/internal/services"
	"github.com/loki4514/parking-lot/internal/storage"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	parkingLot := storage.CreateParkingLot()

	for {
		inputCommand := getInitialCommand()

		if inputCommand == "enter" {
			vehicleType := getVehicleTypeFromUser()
			if vehicleType == storage.Unknown {
				continue
			}

			ticketID, err := OrchestratedEntryTicket(vehicleType, &parkingLot)
			if err != nil {
				fmt.Println("Error creating entry ticket:", err)
				continue
			}

			fmt.Println("✅ Entry ticket created successfully! Ticket ID:", ticketID)

		} else if inputCommand == "exit" {
			exitTicketFlow(&parkingLot)

		} else if inputCommand == "q" {
			fmt.Println("👋 Goodbye, quitting program!")
			break
		}

		time.Sleep(1 * time.Second)
	}
}

// ----------------------------
// Get initial command
// ----------------------------
func getInitialCommand() string {
	fmt.Println("\nAvailable Commands:")
	fmt.Println("ENTER → Create Entry Ticket")
	fmt.Println("EXIT  → Exit Parking Lot")
	fmt.Println("Q     → Quit")

	var input string
	fmt.Print("Enter your command: ")
	fmt.Scanln(&input)
	input = strings.ToLower(strings.TrimSpace(input))

	switch input {
	case "enter":
		return "enter"
	case "exit":
		return "exit"
	case "q":
		return "q"
	default:
		fmt.Println("⚠️ Invalid command, defaulting to 'enter'")
		return "enter"
	}
}

// ----------------------------
// Get vehicle type from user
// ----------------------------
func getVehicleTypeFromUser() storage.VehicleType {
	fmt.Println("\nAvailable Vehicle Types:")
	fmt.Println("Car, Bike, Bicycle, Van")

	var input string
	fmt.Print("Enter your vehicle type: ")
	fmt.Scanln(&input)
	input = strings.ToLower(strings.TrimSpace(input))

	switch input {
	case "car":
		return storage.Car
	case "bike":
		return storage.Bike
	case "bicycle":
		return storage.Bicycle
	case "van":
		return storage.Van
	default:
		fmt.Println("⚠️ Invalid vehicle type. Try again.")
		return storage.Unknown
	}
}

// ----------------------------
// Exit Ticket Flow
// ----------------------------
func exitTicketFlow(parkingLot *storage.ParkingLot) {
	var ticketID string

	for attempts := 0; attempts < 3; attempts++ {
		fmt.Print("Enter your Ticket ID: ")
		fmt.Scanln(&ticketID)

		floorId, spotId, vehicleType, entryTime, err := services.FindTicket(ticketID)
		if err != nil {
			fmt.Println("❌ Invalid Ticket ID. Please try again.")
			continue
		}

		exitTime := time.Now()
		fee := StrategyPattern(vehicleType, entryTime, exitTime)
		fmt.Printf("💰 Parking fee: ₹%d\n", fee)

		if !collectPayment() {
			fmt.Println("❌ Payment failed multiple times. Please contact support.")
			return
		}

		services.UpdateTicket(ticketID, exitTime, fee)
		services.UpdatingSpot(floorId, spotId, vehicleType, parkingLot)

		fmt.Println("✅ Exit successful! Thank you for visiting.")
		return
	}

	fmt.Println("⚠️ Multiple invalid attempts. Please move your vehicle aside and contact support.")
}

// ----------------------------
// Simulate random exit time
// ----------------------------
func simulateExitTime(entryTime time.Time) time.Time {
	randHours := [5]int{1, 2, 3, 4, 5}
	someHour := rand.Intn(len(randHours))
	return time.Now().Add(time.Duration(randHours[someHour]) * time.Hour)
}

// ----------------------------
// Handle payment input
// ----------------------------
func collectPayment() bool {
	for {
		fmt.Print("Payment method (CASH / CARD): ")
		var method string
		fmt.Scanln(&method)

		method = strings.ToUpper(strings.TrimSpace(method))
		if method == "CASH" || method == "CARD" {
			fmt.Printf("✅ Payment successful via %s.\n", method)
			return true
		}

		fmt.Println("⚠️ Invalid payment method. Try again.")
	}
}

// ----------------------------
// Strategy pattern for fee calculation
// ----------------------------
func StrategyPattern(vehicleType storage.VehicleType, entryTime time.Time, exitTime time.Time) int {
	feeAlgo := services.FeeCalculationAlgo{}

	switch vehicleType {
	case storage.Bicycle:
		feeAlgo.SetAlgo(&services.BicycleFeeCalculation{})
	case storage.Bike:
		feeAlgo.SetAlgo(&services.BikeFeeCalculation{})
	case storage.Car:
		feeAlgo.SetAlgo(&services.CarFeeCalculation{})
	case storage.Van:
		feeAlgo.SetAlgo(&services.VanFeeCalculation{})
	default:
		fmt.Println("Unknown vehicle type, setting default ₹100/hr")
		feeAlgo.SetAlgo(&services.CarFeeCalculation{})
	}

	return feeAlgo.Calculate(entryTime, exitTime)
}

// ----------------------------
// Orchestrates ticket creation
// ----------------------------
func OrchestratedEntryTicket(vehicleType storage.VehicleType, parkingLot *storage.ParkingLot) (string, error) {
	entryTicket, err := services.CreateEntryTicket(vehicleType, parkingLot)
	if err != nil {
		return "", err
	}
	return entryTicket.TicketID, nil
}
