package report

import "store/catalog"

// TotalValue считает суммарную стоимость склада: Price * Stock по всем товарам.
func TotalValue(items []catalog.Item) float64 {
	var total float64
	for _, item := range items {
		total += item.Price * float64(item.Stock)
	}
	return total
}
