package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Owner   string
	balance float64
}

func (a *BankAccount) Deposit(amount float64) {
	a.balance += amount
}

func (a *BankAccount) Withdraw(amount float64) error {
	if amount > a.balance {
		return errors.New("недостаточно средств")
	}
	a.balance -= amount
	return nil
}

func Transfer(from, to *BankAccount, amount float64) error {
	if err := from.Withdraw(amount); err != nil {
		return fmt.Errorf("перевод не выполнен: %w", err)
	}
	to.Deposit(amount)
	return nil
}

func main() {
	alice := &BankAccount{Owner: "Alice", balance: 1000}
	bob := &BankAccount{Owner: "Bob", balance: 200}

	if err := Transfer(alice, bob, 300); err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Printf("Alice: %.2f, Bob: %.2f\n", alice.balance, bob.balance)

	if err := Transfer(bob, alice, 10000); err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Printf("Alice: %.2f, Bob: %.2f\n", alice.balance, bob.balance)
}
