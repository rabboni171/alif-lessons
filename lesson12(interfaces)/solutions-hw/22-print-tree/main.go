package main

import (
	"fmt"
	"strings"
)

type Category struct {
	Name string
	Sub  []Category
}

func PrintTree(v any, depth int) {
	switch x := v.(type) {
	case Category:
		fmt.Println(strings.Repeat("  ", depth) + x.Name)
		PrintTree(x.Sub, depth+1)
	case []Category:
		for _, category := range x {
			PrintTree(category, depth)
		}
	}
}

func main() {
	root := Category{
		Name: "Электроника",
		Sub: []Category{
			{
				Name: "Телефоны",
				Sub: []Category{
					{Name: "Смартфоны"},
					{Name: "Кнопочные"},
				},
			},
			{
				Name: "Ноутбуки",
				Sub: []Category{
					{Name: "Игровые"},
					{Name: "Офисные"},
				},
			},
		},
	}

	PrintTree(root, 0)
}
