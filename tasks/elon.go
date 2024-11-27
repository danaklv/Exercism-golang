package tasks

import "fmt"
// TODO: define the 'Drive()' method
func (car *Car) Drive() {
    if car.battery >= car.batteryDrain {
        car.battery -= car.batteryDrain
        car.distance += car.speed
    }
}

// TODO: define the 'DisplayDistance() string' method
func (car Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", car.distance)
}


// TODO: define the 'DisplayBattery() string' method

func (car Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", car.battery)
}


// TODO: define the 'CanFinish(trackDistance int) bool' method

func (car Car) CanFinish(trackDistance int) bool {
    for car.battery >= car.batteryDrain {
        trackDistance -= car.speed 
        car.battery -= car.batteryDrain
    }
    return trackDistance <= 0
}


// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int

	battery  int
	distance int
}

// NewCar creates a new car with given specifications.
func NewCar(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}