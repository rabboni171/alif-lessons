package main

import (
	"fmt"

	"importcycle/a"
	"importcycle/b"
)

func main() {
	fmt.Println(a.Hello())
	fmt.Println(b.Hello())
}

// До исправления пакет a импортировал b, а пакет b импортировал a
// напрямую (без common). `go build ./...` в таком виде падал с ошибкой:
//
//   package importcycle/a
//   	imports importcycle/b from a.go
//   	imports importcycle/a from b.go: import cycle not allowed
//
// Исправление: то общее, что было нужно обоим (Greeting), вынесено в
// третий пакет common, от которого зависят и a, и b, а сам common ни
// от одного из них не зависит - цикл разорван.
