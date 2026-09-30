package main

import "fmt"

func Withdraw(balance, amount float64) (float64, error) {
	if amount < 0 {
		return balance, fmt.Errorf("сумма снятия не может быть отрицательной: %.2f", amount)
	}
	if amount > balance {
		return balance, fmt.Errorf("недостаточно средств: на балансе %.2f, а нужно снять %.2f", balance, amount)
	}
	return balance - amount, nil
}

func main() {
	balance := 1000.0

	newBalance, err := Withdraw(balance, 300)
	if err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Printf("успешное снятие, новый баланс: %.2f\n", newBalance)
	}

	_, err = Withdraw(balance, -50)
	if err != nil {
		fmt.Println("ошибка:", err)
	}

	_, err = Withdraw(balance, 5000)
	if err != nil {
		fmt.Println("ошибка:", err)
	}
}
