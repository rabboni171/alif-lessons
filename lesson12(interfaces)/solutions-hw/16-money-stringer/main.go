package main

import "fmt"

type Money struct {
	Amount   float64
	Currency string
}

func (m Money) String() string {
	return fmt.Sprintf("%.2f %s", m.Amount, m.Currency)
}

func main() {
	prices := []Money{
		{Amount: 150, Currency: "USD"},
		{Amount: 2500, Currency: "UZS"},
		{Amount: 99.9, Currency: "EUR"},
	}

	for _, p := range prices {
		fmt.Println(p)
	}
}
