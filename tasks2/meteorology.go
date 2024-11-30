package tasks2

import "strconv"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (t TemperatureUnit) String() string {
	if t == 0 {
		return "°C"
	} 
	return "°F"
}



type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type

func (t Temperature) String() string {
  return  strconv.Itoa(t.degree)  + " " +	t.unit.String()

}


// Add a String method to Speed
type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit

func (s SpeedUnit) String() string {
	if s == 0 {
		return "km/h"
	}
	return "mph"
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
	return strconv.Itoa(s.magnitude) + " " + s.unit.String()
}



type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData

func (m MeteorologyData) String() string {
	h := strconv.Itoa(m.humidity)
	return m.location + ": " + m.temperature.String() + ", Wind " + m.windDirection + " at " + m.windSpeed.String() + ", " + h + "% Humidity" 
	//return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity",md.location, md.temperature, md.windDirection,md.windSpeed,md.humidity)
}
