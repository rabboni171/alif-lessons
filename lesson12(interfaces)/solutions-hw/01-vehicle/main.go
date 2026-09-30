package main

import "fmt"

type Vehicle interface {
	Start() string
	Stop() string
}

type Car struct {
	Name string
}

func (c Car) Start() string {
	return "Car " + c.Name + " заводится"
}

func (c Car) Stop() string {
	return "Car " + c.Name + " глохнет"
}

type Bike struct {
	Name string
}

func (b Bike) Start() string {
	return "Bike " + b.Name + " заводится"
}

func (b Bike) Stop() string {
	return "Bike " + b.Name + " глохнет"
}

type Bus struct {
	Name string
}

func (b Bus) Start() string {
	return "Bus " + b.Name + " заводится"
}

func (b Bus) Stop() string {
	return "Bus " + b.Name + " глохнет"
}

func main() {
	vehicles := []Vehicle{
		Car{Name: "Toyota"},
		Bike{Name: "Stels"},
		Bus{Name: "MAN"},
	}

	for _, v := range vehicles {
		fmt.Println(v.Start())
		fmt.Println(v.Stop())
	}
}
