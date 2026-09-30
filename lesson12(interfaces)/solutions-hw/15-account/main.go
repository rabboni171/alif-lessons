package main

import (
	"errors"
	"fmt"
)

type Depositor interface {
	Deposit(amount float64) error
}

type Withdrawer interface {
	Withdraw(amount float64) error
}

type Account interface {
	Depositor
	Withdrawer
}

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
	if amount < 0 {
		return errors.New("сумма снятия не может быть отрицательной")
	}
	if amount > w.balance {
		return errors.New("недостаточно денег на счёте")
	}
	w.balance -= amount
	return nil
}

func main() {
	var account Account = &Wallet{}

	if err := account.Deposit(1000); err != nil {
		fmt.Println("Ошибка:", err)
	}

	if err := account.Withdraw(300); err != nil {
		fmt.Println("Ошибка:", err)
	}

	if err := account.Withdraw(10000); err != nil {
		fmt.Println("Ошибка:", err)
	}

	wallet := account.(*Wallet)
	fmt.Printf("Итоговый баланс: %.2f\n", wallet.balance)
}
