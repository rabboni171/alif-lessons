package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Invoice struct {
	Client string
	Amount float64
}

func (i Invoice) String() string {
	return fmt.Sprintf("%s: %.2f", i.Client, i.Amount)
}

func WriteReport(w io.Writer, invoices []Invoice) float64 {
	total := 0.0
	for _, invoice := range invoices {
		fmt.Fprintln(w, invoice)
		total += invoice.Amount
	}
	return total
}

func main() {
	invoices := []Invoice{
		{Client: "Азиз", Amount: 1500},
		{Client: "Диёр", Amount: 2300},
		{Client: "Фарход", Amount: 800},
	}

	totalStdout := WriteReport(os.Stdout, invoices)
	fmt.Printf("Общая сумма: %.2f\n", totalStdout)

	var builder strings.Builder
	totalBuilder := WriteReport(&builder, invoices)
	fmt.Println(builder.String())
	fmt.Printf("Общая сумма: %.2f\n", totalBuilder)
}
