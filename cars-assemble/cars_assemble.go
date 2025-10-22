package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100.0
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
//
//nolint:gosec // G115: integer overflow isn't possible.
func CalculateCost(carsCount int) uint {
	if carsCount < 0 {
		panic("carsCount must not be negative")
	}
	numGroupsOf10 := uint(carsCount / 10)
	numRemaining := uint(carsCount % 10)
	return numGroupsOf10*95000 + numRemaining*10000
}
