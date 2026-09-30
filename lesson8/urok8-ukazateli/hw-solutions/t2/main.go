// Задача 2: изменение температуры через указатель.
package main

import "fmt"

func main() {
	temperature := 18.5
	p := &temperature

	fmt.Println("До:   ", temperature)
	*p = *p - 5
	fmt.Println("После:", temperature)
}
