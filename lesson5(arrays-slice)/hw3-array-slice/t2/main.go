package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}

	slice := arr[1:4] // со 2-го по 4-й элемент (индексы 1,2,3)

	fmt.Println("срез:", slice)
	fmt.Println("len:", len(slice))
	fmt.Println("cap:", cap(slice)) // cap считает от начала среза и до конца исходного массива
}
