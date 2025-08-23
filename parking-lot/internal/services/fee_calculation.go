package services

import (
	"math"
	"time"
)

// FeeCalculation strategy interface
type FeeCalculation interface {
	CalculateFee(entryTime time.Time) int
}

// Common fare calculation logic
func estimateFare(basePay int, entryTime time.Time) int {
	exitTime := time.Now()
	diff := exitTime.Sub(entryTime)

	hours := int(math.Abs(diff.Hours()))
	minutes := int(math.Abs(diff.Minutes())) % 60

	if hours <= 0 {
		return basePay
	}

	calculatedFee := hours * basePay
	if minutes >= 15 {
		calculatedFee += basePay
	}
	return calculatedFee
}

// ----------------------------
// Bicycle Fee Calculation
// ----------------------------
type BicycleFeeCalculation struct{}

func (b *BicycleFeeCalculation) CalculateFee(entryTime time.Time) int {
	return estimateFare(20, entryTime) // ₹20 per hour
}

// ----------------------------
// Bike Fee Calculation
// ----------------------------
type BikeFeeCalculation struct{}

func (b *BikeFeeCalculation) CalculateFee(entryTime time.Time) int {
	return estimateFare(50, entryTime) // ₹50 per hour
}

// ----------------------------
// Car Fee Calculation
// ----------------------------
type CarFeeCalculation struct{}

func (c *CarFeeCalculation) CalculateFee(entryTime time.Time) int {
	return estimateFare(100, entryTime) // ₹100 per hour
}

// ----------------------------
// Van Fee Calculation
// ----------------------------
type VanFeeCalculation struct{}

func (v *VanFeeCalculation) CalculateFee(entryTime time.Time) int {
	return estimateFare(150, entryTime) // ₹150 per hour
}

type FeeCalculationAlgo struct {
	feeAlgo FeeCalculation
}

func (feeCalculation *FeeCalculationAlgo) SetAlgo(algo FeeCalculation) {
	feeCalculation.feeAlgo = algo

}
