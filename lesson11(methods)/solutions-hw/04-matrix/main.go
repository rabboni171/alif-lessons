package main

import "fmt"

type Matrix [][]float64

func (m Matrix) Rows() int {
	return len(m)
}

func (m Matrix) Cols() int {
	if len(m) == 0 {
		return 0
	}
	return len(m[0])
}

func (m Matrix) Sum() float64 {
	total := 0.0
	for _, row := range m {
		for _, v := range row {
			total += v
		}
	}
	return total
}

func (m Matrix) Transpose() Matrix {
	result := make(Matrix, m.Cols())
	for i := range result {
		result[i] = make([]float64, m.Rows())
	}

	for r, row := range m {
		for c, v := range row {
			result[c][r] = v
		}
	}
	return result
}

func main() {
	m := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("Строк:", m.Rows())
	fmt.Println("Столбцов:", m.Cols())
	fmt.Println("Сумма элементов:", m.Sum())

	t := m.Transpose()
	fmt.Println("Транспонированная матрица:")
	for _, row := range t {
		fmt.Println(row)
	}
}
