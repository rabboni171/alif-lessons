package catalog

// Product - товар в магазине.
type Product struct {
	Name  string
	Price float64
}

// TotalPrice суммирует цены всех товаров в срезе.
func TotalPrice(products []Product) float64 {
	var total float64
	for _, p := range products {
		total += p.Price
	}
	return total
}
