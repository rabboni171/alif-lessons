package main

import (
	"fmt"
	"sort"

	"registrydemo/registry"
	_ "registrydemo/reverse"
	_ "registrydemo/upper"
)

func main() {
	text := "Привет, мир"

	handlers := registry.All()

	names := make([]string, 0, len(handlers))
	for name := range handlers {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s: %s\n", name, handlers[name].Handle(text))
	}
}
