package main

import "fmt"

func sayHello() {
	fmt.Println("Hello1")
}

/* А что если функция нужна только один раз?
1. Посчитать что-нибудь.
2. Вывести результат.
3. Больше эта функция никогда не понадобится.
*/

// для этого есть анонимные функции

func main() {
	//sayHello()

	func() {
		fmt.Println("Hello from anonymous func")
	}()

	func() {
		fmt.Println("Я выполняюсь только один раз")
	}()

	fmt.Println("Конец программы")

	// Можно передавать параметры
	func(name string) {
		fmt.Println("Привет,", name)
	}("Akbar")

	// Можно вернуть значение
	result := func(a, b int) int {
		return a + b
	}(5, 3)

	fmt.Println(result)

	// Можно сохранить функцию в переменную
	sum := func(a, b int) int {
		return a + b
	}

	fmt.Println(sum(10, 20))
}

/*

func() {
		fmt.Println("Hello from anonymous func")
	}()

Последние круглые скобки () означают вызови функцию прямо сейчас
*/
