// Задача 15: рекорд спортсмена - обновление лучшего результата через указатель.
package main

import "fmt"

func trackRecord(best *float64, attempt float64) bool {
	if attempt > *best {
		*best = attempt
		return true
	}
	return false
}

func main() {
	attempts := []float64{5.2, 5.8, 5.5, 6.1, 5.9, 6.3}

	var best float64
	for i, attempt := range attempts {
		if trackRecord(&best, attempt) {
			fmt.Printf("Попытка %d: %.1f - новый рекорд!\n", i+1, attempt)
		} else {
			fmt.Printf("Попытка %d: %.1f\n", i+1, attempt)
		}
	}

	fmt.Printf("Итоговый рекорд: %.1f\n", best)
}
