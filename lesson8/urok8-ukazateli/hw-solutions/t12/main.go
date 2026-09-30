// Задача 12: подсчёт голосов - три счётчика меняются через указатели.
package main

import "fmt"

func tallyVote(candidate string, countA, countB, countC *int) {
	switch candidate {
	case "Алиев":
		*countA++
	case "Борисов":
		*countB++
	case "Валиева":
		*countC++
	}
}

func main() {
	votes := []string{
		"Алиев", "Борисов", "Алиев", "Валиева", "Борисов",
		"Алиев", "Алиев", "Валиева", "Борисов", "Алиев",
	}

	var countA, countB, countC int
	for _, name := range votes {
		tallyVote(name, &countA, &countB, &countC)
	}

	fmt.Println("Алиев:  ", countA)
	fmt.Println("Борисов:", countB)
	fmt.Println("Валиева:", countC)
}
