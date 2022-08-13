// Package weather implements weather related functions.
package weather

// CurrentCondition represents current weather condition.
var CurrentCondition string

// CurrentLocation represents the given location.
var CurrentLocation string

// Forecast returns string representation of weather for given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
