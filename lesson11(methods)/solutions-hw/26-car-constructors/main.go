package main

import "fmt"

type Car struct {
	Brand, Model string
	mileage      int
}

func NewCar(brand, model string) *Car {
	return &Car{Brand: brand, Model: model, mileage: 0}
}

func NewUsedCar(brand, model string, mileage int) *Car {
	return &Car{Brand: brand, Model: model, mileage: mileage}
}

func (c *Car) Drive(km int) {
	c.mileage += km
}

func main() {
	newCar := NewCar("Toyota", "Camry")
	usedCar := NewUsedCar("Honda", "Civic", 45000)

	newCar.Drive(120)
	usedCar.Drive(80)

	fmt.Printf("%s %s: пробег %d км\n", newCar.Brand, newCar.Model, newCar.mileage)
	fmt.Printf("%s %s: пробег %d км\n", usedCar.Brand, usedCar.Model, usedCar.mileage)
}
