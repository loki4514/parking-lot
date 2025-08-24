package services

import (
	"time"
)

// ----------------------------
// FeeCalculation Strategy Interface
// ----------------------------
type FeeCalculation interface {
	CalculateFee(entryTime, exitTime time.Time) int
}

// ----------------------------
// Common fare calculation logic
// ----------------------------
func estimateFare(basePay int, entryTime, exitTime time.Time) int {
	if exitTime.Before(entryTime) {
		exitTime = time.Now()
	}

	diff := exitTime.Sub(entryTime)
	hours := int(diff.Hours())
	minutes := int(diff.Minutes()) % 60

	// Minimum 1 hour charge
	if hours <= 0 {
		return basePay
	}

	calculatedFee := hours * basePay

	// If leftover minutes >= 15, charge for 1 extra hour
	if minutes >= 15 {
		calculatedFee += basePay
	}

	return calculatedFee
}

// ----------------------------
// Bicycle Fee Calculation
// ----------------------------
type BicycleFeeCalculation struct{}

func (b *BicycleFeeCalculation) CalculateFee(entryTime, exitTime time.Time) int {
	return estimateFare(20, entryTime, exitTime)
}

// ----------------------------
// Bike Fee Calculation
// ----------------------------
type BikeFeeCalculation struct{}

func (b *BikeFeeCalculation) CalculateFee(entryTime, exitTime time.Time) int {
	return estimateFare(50, entryTime, exitTime)
}

// ----------------------------
// Car Fee Calculation
// ----------------------------
type CarFeeCalculation struct{}

func (c *CarFeeCalculation) CalculateFee(entryTime, exitTime time.Time) int {
	return estimateFare(100, entryTime, exitTime)
}

// ----------------------------
// Van Fee Calculation
// ----------------------------
type VanFeeCalculation struct{}

func (v *VanFeeCalculation) CalculateFee(entryTime, exitTime time.Time) int {
	return estimateFare(150, entryTime, exitTime)
}

// ----------------------------
// Context for Strategy
// ----------------------------
type FeeCalculationAlgo struct {
	feeAlgo FeeCalculation
}

func (f *FeeCalculationAlgo) SetAlgo(algo FeeCalculation) {
	f.feeAlgo = algo
}

func (f *FeeCalculationAlgo) Calculate(entryTime, exitTime time.Time) int {
	if f.feeAlgo == nil {
		return 0
	}
	return f.feeAlgo.CalculateFee(entryTime, exitTime)
}
