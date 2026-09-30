package main

import (
	"fmt"
)

func main() {
	var x, y float64

	fmt.Println("Введите координаты X и Y (через пробел, 0 нельзя)")
	fmt.Scan(&x, &y)

	if x == 0 || y == 0 {
		fmt.Println("Не балуйтесь")
		return
	}

	switch {
	case x > 0 && y > 0:
		fmt.Println("1 четверть")
	case x < 0 && y > 0:
		fmt.Println("2 четверть")
	case x < 0 && y < 0:
		fmt.Println("3 четверть")
	case x > 0 && y < 0:
		fmt.Println("4 четверть")
	}
}
