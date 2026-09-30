// Задача 1: значение, адрес и тип адреса переменной.
package main

import "fmt"

func main() {
	age := 25
	fmt.Println("Значение age:", age)
	fmt.Println("Адрес age:  ", &age)
	fmt.Printf("Тип &age:    %T\n", &age)
}
