package main

import "fmt"

type Discount interface {
	Apply(price float64) float64
}

type PercentDiscount struct {
	Percent float64
}

func (d PercentDiscount) Apply(price float64) float64 {
	return price - price*d.Percent/100
}

type FixedDiscount struct {
	Amount float64
}

func (d FixedDiscount) Apply(price float64) float64 {
	return price - d.Amount
}

type NoDiscount struct{}

func (d NoDiscount) Apply(price float64) float64 {
	return price
}

func main() {
	price := 1000.0

	discounts := []Discount{
		PercentDiscount{Percent: 10},
		FixedDiscount{Amount: 50},
		NoDiscount{},
		PercentDiscount{Percent: 5},
	}

	for _, d := range discounts {
		price = d.Apply(price)
		fmt.Printf("После применения скидки: %.2f\n", price)
	}
}
