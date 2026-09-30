package main

import "fmt"

func main() {
	count := 0

	increment := func() {
		count++
		fmt.Println(count)
	}

	//функция замкнула вокруг себя внешнюю переменную.
	/*
		она не копирует переменную,не хранит ссылку на нее,
		поэтому значение меняется.
	*/

	increment()
	increment()
	increment()

	name := "Ali"

	printName := func() {
		fmt.Println(name)
	}

	name = "Bob"

	printName()

	x := 10

	if x > 5 {
		func() {
			fmt.Println("x больше 5")
		}()
	}

	//hello()

	execute(func() {
		fmt.Println("Привет!")
	})
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

func countdown(n int) {
	//n=5 
	if n == 0 {
		fmt.Println("Стоп")
		return
	}

	fmt.Println(n) // 5 4 3 2 1 

	countdown(n - 1) //countdown(4) countdown(3) 
}

func hello() {
	func() {
		fmt.Println("Я внутри hello()")
	}()
}

func getPrinter() func() {
	return func() {
		fmt.Println("Hello")
	}
}

func execute(f func()) {
	f()
}

// что выведет?
func hello1() {
	hello1()
}
