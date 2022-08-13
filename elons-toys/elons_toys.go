package elon

import "fmt"

func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d", car.battery) + "%"
}

func (car Car) CanFinish(trackDistance int) bool {
	startDistance := car.distance
	for car.battery > 0 && car.battery >= car.batteryDrain {
		car.Drive()
		if (car.distance - startDistance) >= trackDistance {
			return true
		}
	}
	return false
}
