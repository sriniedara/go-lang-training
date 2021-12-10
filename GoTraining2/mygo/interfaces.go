package mygo

import "fmt"

func TestInterfaces() {

	var vehicle Vehicle

	var car Car
	car.name = "DZire"
	car.isMoving = false
	vehicle = &car

	vehicle.accelerate(10)
	vehicle.describe()

	//vehicle = &Bike{name: "Suzuki Access 125", isMoving: false}
}

type Vehicle interface {
	applyBrake()
	accelerate(by uint8) uint8
	changeGear(to uint8) uint8
	describe()
}

type Car struct {
	name string
	VehicleDetails
}

type Bike struct {
	name string
	VehicleDetails
}

type VehicleDetails struct {
	isMoving bool
	speed    uint8
	gear     uint8
}

func (car *Car) applyBrake() {
	car.isMoving = false
}

func (bike *Bike) applyBrake() {
	bike.isMoving = false
}

func (car *Car) accelerate(by uint8) uint8 {
	car.isMoving = true
	car.speed = car.speed + by
	return car.speed
}

func (bike *Bike) accelerate(by uint8) uint8 {
	bike.isMoving = true
	bike.speed = bike.speed + by
	return bike.speed
}

func (car *Car) changeGear(to uint8) uint8 {
	current := car.gear
	car.gear = to
	return current
}

func (bike *Bike) changeGear(to uint8) uint8 {
	current := bike.gear
	bike.gear = to
	return current
}

func (car *Car) describe() {
	fmt.Println("Car Details: Name:", car.name, ",Gear:", car.gear, ",Speed:", car.speed,
		"isMoving:", car.isMoving)
}

func (bike *Bike) describe() {
	fmt.Println("Bike Details: Name:", bike.name, ",Gear:", bike.gear, ",Speed:", bike.speed,
		"isMoving:", bike.isMoving)
}
