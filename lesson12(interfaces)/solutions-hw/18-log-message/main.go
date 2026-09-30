package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func LogMessage(w io.Writer, msg string) {
	fmt.Fprintln(w, msg)
}

func main() {
	LogMessage(os.Stdout, "Сообщение в консоль")

	var builder strings.Builder
	LogMessage(&builder, "Сообщение в билдер")

	fmt.Println(builder.String())
}
