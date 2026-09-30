package main

import "fmt"

type Category struct {
	Name string
	Sub  []Category
}

func CountAll(c Category) int {
	count := 1
	for _, sub := range c.Sub {
		count += CountAll(sub)
	}
	return count
}

func main() {
	electronics := Category{
		Name: "Электроника",
		Sub: []Category{
			{Name: "Телефоны"},
			{
				Name: "Компьютеры",
				Sub: []Category{
					{Name: "Ноутбуки"},
					{Name: "Настольные ПК"},
				},
			},
		},
	}

	fmt.Println("Всего категорий:", CountAll(electronics))
}
