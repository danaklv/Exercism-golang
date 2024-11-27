// Package weather provides tools to represent condition on current location.
package tasks

// CurrentCondition represent a current condition.
var CurrentCondition string
// CurrentLocation represent a current location.
var CurrentLocation string

// Forecast takes city and condition arguments as string, and return a string with condition at this location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}