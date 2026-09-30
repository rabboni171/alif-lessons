package main

import "fmt"

func main() {
	var words []string // пустой срез строк

	words = append(words, "go")
	fmt.Println(words, "len:", len(words))

	words = append(words, "это")
	fmt.Println(words, "len:", len(words))

	words = append(words, "просто")
	fmt.Println(words, "len:", len(words))
}
