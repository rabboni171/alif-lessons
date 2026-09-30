package main

import (
	"fmt"

	"wallet/wallet"
)

func main() {
	w := wallet.Wallet{}

	// w.balance напрямую не скомпилируется - поле не экспортируется
	// из пакета wallet: "w.balance undefined (cannot refer to unexported
	// field or method balance)".
	// fmt.Println(w.balance)

	w.Deposit(100)
	w.Deposit(50)

	fmt.Println("Баланс:", w.Balance())
}
