// Package weather provides tools to forecast the current
// weather condition of various cities in Goblinocus.
package weather

var (
	// CurrentCondition represents the current weather condition.
	//nolint:gochecknoglobals
	CurrentCondition string
	// CurrentLocation represents the city.
	//nolint:gochecknoglobals
	CurrentLocation string
)

// Forecast returns a description of the current weather condition for the given city.
// The 'city' parameter represents the city.
// The 'condition' parameter represents the current weather condition.
// It returns a string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
