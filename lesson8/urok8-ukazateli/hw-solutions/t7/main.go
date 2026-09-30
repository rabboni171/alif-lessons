// Задача 7: swap для строк через указатели.
package main

import "fmt"

func swapStrings(a, b *string) {
	*a, *b = *b, *a
}

func main() {
	student1, student2 := "Али", "Вера"
	fmt.Println("До:   ", student1, student2)

	swapStrings(&student1, &student2)
	fmt.Println("После:", student1, student2)
}
