// Задача 5: удалить закончившийся товар и проверить через comma-ok.
package main

import "fmt"

func main() {
	stock := map[string]int{
		"хлеб":   10,
		"молоко": 0,
		"яйца":   15,
	}

	delete(stock, "молоко")

	if _, ok := stock["молоко"]; !ok {
		fmt.Println("молоко закончилось, товар удалён со склада")
	}

	fmt.Println(stock)
}
