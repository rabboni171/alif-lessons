package main

import "fmt"

func main() {
	var ( 
		number1 int
		number2 int
		number3 int
		count int
	)

	fmt.Println("Введите первое число:")
	fmt.Scan(&number1)

	fmt.Println("Введите второе число:")
	fmt.Scan(&number2)

	fmt.Println("Введите третье число:")
	fmt.Scan(&number3)

	if number1 > 0 {
		count += 1
	} 

	if number2 > 0 {
		count += 1
	} 

	if number3 > 0 {
		count += 1
	} 

	fmt.Println("Количество положительных чисел:", count)
}
