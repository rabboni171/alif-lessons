// Задача 3: изменение строки через указатель.
package main

import "fmt"

func main() {
	city := "Ташкент"
	p := &city

	fmt.Println("До:   ", city)
	*p = "Самарканд"
	fmt.Println("После:", city)
}
