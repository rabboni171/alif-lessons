package main

import "fmt"

func main() {
	var (
		balance float64
		sum     int
	)

	fmt.Print("Укажите ваш баланс: ")
	fmt.Scan(&balance)

	// До(рефакторинга)
	// if balance == 0 {
	// 	fmt.Println("Вам нечего снимать со счёта")
	// 	return
	// } else if balance < 0 {
	// 	fmt.Println("Пожалуйста укажите корректную сумму на вашем счёте")
	// 	return
	// }

	//После рефакторинга
	if balance <= 0 {
		fmt.Println("Укажите корректную сумму на вашем счёте")
		return
	}

	fmt.Print("Укажите сумму для снятия: ")
	fmt.Scan(&sum)

	floatedSum := float64(sum)

	if sum <= 0 {
		fmt.Println("Ошибка: некорректная сумма")
		return
	} else if floatedSum > balance {
		fmt.Println("Недостаточно средств")
	} else {
		balance -= floatedSum
	}

	fmt.Printf("Сумма снятия: %d\n", sum)
	fmt.Printf("Остаток на счёте: %v\n", balance)
}
