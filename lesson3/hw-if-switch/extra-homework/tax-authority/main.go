package main

import "fmt"

func main() {
	
	const (
		lowTax    = 5
		middleTax = 10
		highTax   = 15
	)

	var tax,salary float64

	fmt.Print("Укажите вашу заработную плату: ")
	fmt.Scan(&salary)

	if salary < 1000 {
		fmt.Println("Заработная плата не может быть менее 1000 сомони")
		return
	}

	if salary <= 3000 {
		tax = lowTax
	} else if salary <= 10000 {
		tax = middleTax
	} else {
		tax = highTax
	}

	salary = salary * (100 - tax) / 100

	fmt.Printf("Ваш подоходный налог составляет %v%%\n", tax)
	fmt.Printf("Ваша заработная плата с учётом налога составляет: %.2f\n", salary)
}
