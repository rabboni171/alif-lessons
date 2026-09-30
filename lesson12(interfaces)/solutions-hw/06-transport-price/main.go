package main

import "fmt"

type Transport interface {
	Price(distanceKm float64) float64
}

type Taxi struct{}

func (t Taxi) Price(distanceKm float64) float64 {
	return 10 + distanceKm*3
}

type Bus struct{}

func (b Bus) Price(distanceKm float64) float64 {
	return 5
}

type Bike struct{}

func (b Bike) Price(distanceKm float64) float64 {
	return distanceKm * 1.5
}

func main() {
	distance := 8.0

	transports := map[string]Transport{
		"Taxi": Taxi{},
		"Bus":  Bus{},
		"Bike": Bike{},
	}

	cheapestName := ""
	cheapestPrice := 0.0

	for name, t := range transports {
		price := t.Price(distance)
		fmt.Printf("%s: %.2f\n", name, price)

		if cheapestName == "" || price < cheapestPrice {
			cheapestName = name
			cheapestPrice = price
		}
	}

	fmt.Printf("Дешевле всего: %s (%.2f)\n", cheapestName, cheapestPrice)
}
