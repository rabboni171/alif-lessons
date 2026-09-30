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
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if amount > w.balance {
		return errors.New("insufficient balance")
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() float64 {
	return w.balance
}

func main() {
	wallet := &Wallet{}

	var acc Account = wallet

	if err := acc.Deposit(1000); err != nil {
		fmt.Println(err)
	}

	if err := acc.Withdraw(250); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Balance: %.2f\n", wallet.Balance())
}
