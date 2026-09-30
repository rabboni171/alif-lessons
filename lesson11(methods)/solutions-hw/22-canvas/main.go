package main

import (
	"fmt"
)

type Canvas struct {
	grid [][]rune
}

func NewCanvas(width, height int) *Canvas {
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}
	return &Canvas{grid: grid}
}

func (c *Canvas) SetPixel(x, y int, ch rune) error {
	if y < 0 || y >= len(c.grid) || x < 0 || x >= len(c.grid[0]) {
		return fmt.Errorf("координаты (%d, %d) вне холста", x, y)
	}
	c.grid[y][x] = ch
	return nil
}

func (c Canvas) Print() {
	for _, row := range c.grid {
		fmt.Println(string(row))
	}
}

func main() {
	canvas := NewCanvas(8, 4)

	for x := 0; x < 8; x++ {
		canvas.SetPixel(x, 1, '*')
	}

	if err := canvas.SetPixel(20, 20, '*'); err != nil {
		fmt.Println("Ошибка:", err)
	}

	canvas.Print()
}
