package speed

type Car struct {
	battery, batteryDrain, speed, distance int
}

type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{speed: speed, batteryDrain: batteryDrain, battery: 100}
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if canDrive(car) {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
	return car
}

func canDrive(car Car) bool {
	return car.battery >= car.batteryDrain
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	for canDrive(car) && car.distance < track.distance {
		// by default, struct is passed by value.
		car = Drive(car)
	}
	return car.distance >= track.distance
}
