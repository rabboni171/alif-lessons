// Задача 10: указатель через new без отдельной переменной.
package main

import "fmt"

func main() {
	p := new(int)
	*p = 100
	fmt.Println(*p)
}
