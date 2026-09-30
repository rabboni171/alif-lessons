package main

import (
	"fmt"

	"store/catalog"
	"store/report"
)

func main() {
	fmt.Println("Общая стоимость склада:", report.TotalValue(catalog.Items))
}
