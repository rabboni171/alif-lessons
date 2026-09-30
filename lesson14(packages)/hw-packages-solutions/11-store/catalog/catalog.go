package catalog

// Item - товар на складе.
type Item struct {
	Name  string
	Price float64
	Stock int
}

// Items - "предзагруженный каталог", заполняется автоматически при
// старте программы через init(), до вызова main().
var Items []Item

func init() {
	Items = []Item{
		{Name: "Тетрадь", Price: 3.5, Stock: 50},
		{Name: "Ручка", Price: 1.2, Stock: 120},
		{Name: "Учебник Go", Price: 25, Stock: 15},
		{Name: "Рюкзак", Price: 40, Stock: 8},
		{Name: "Ластик", Price: 0.8, Stock: 200},
	}
}
