package main

import (
	"fmt"
	"strings"
)

func main() {
	type Fuel string

	const (
		A92    Fuel = "92"
		A95    Fuel = "95"
		A98    Fuel = "98"
		DIESEL Fuel = "diesel"
	)

	var (
		fuelType string
		quantity float64
		price    float64
		total    float64
	)

	fmt.Print("Укажите тип топлива (92, 95, 98, diesel): ")
	fmt.Scan(&fuelType)

	fuelChoose := Fuel(strings.ToLower(fuelType))
	switch fuelChoose {
	case A92:
		price = 10
	case A95:
		price = 12
	case A98:
		price = 15
	case DIESEL:
		price = 11
	default:
		fmt.Println("На нашей АЗС нет такого топлива")
		return
	}

	fmt.Print("Укажите необходимое количество в литрах: ")
	fmt.Scan(&quantity)

	if quantity < 1 {
		fmt.Println("Минимальный заказ от 1 литра")
		return
	}

	total = quantity * price

	fmt.Println("Цена за литр топлива: ", price)
	fmt.Printf("Количество: %.2f литра(ов)\n", quantity)
	fmt.Printf("Итого: %.2f сомони", quantity*price)
}
