package jedlik

import "fmt"

func (car *Car) canDrive() bool {
	return car.battery >= car.batteryDrain
}

func (car *Car) Drive() {
	if car.canDrive() {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car *Car) CanFinish(trackDistance int) bool {
	for car.canDrive() && car.distance < trackDistance {
		car.Drive()
	}
	return car.distance >= trackDistance
}
