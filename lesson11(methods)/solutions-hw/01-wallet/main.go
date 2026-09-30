package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance float64
}

func (w *Wallet) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("сумма пополнения не может быть отрицательной")
	}
	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount > w.balance {
		return fmt.Errorf("не хватает денег: на счету %.2f, а нужно %.2f", w.balance, amount)
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() float64 {
	return w.balance
}

func main() {
	w := &Wallet{}

	w.Deposit(1000)
	fmt.Printf("Баланс после пополнения: %.2f\n", w.Balance())

	w.Withdraw(300)
	fmt.Printf("Баланс после списания: %.2f\n", w.Balance())

	w.Deposit(150)
	fmt.Printf("Баланс после пополнения: %.2f\n", w.Balance())

	if err := w.Withdraw(10000); err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Printf("Баланс не изменился: %.2f\n", w.Balance())
}
