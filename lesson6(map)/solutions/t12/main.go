// Задача 12 (для быстрых): частота слов в тексте + самое частое слово.
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	text := "go простой go быстрый go надёжный язык"

	words := strings.Fields(text)
	counts := make(map[string]int)
	for _, w := range words {
		counts[strings.ToLower(w)]++
	}

	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("Частота слов:")
	for _, k := range keys {
		fmt.Printf("%-10s %d\n", k, counts[k])
	}

	maxWord, maxCount := "", 0
	for w, c := range counts {
		if c > maxCount {
			maxWord, maxCount = w, c
		}
	}
	fmt.Printf("Самое частое слово: %q (%d раз)\n", maxWord, maxCount)
}
