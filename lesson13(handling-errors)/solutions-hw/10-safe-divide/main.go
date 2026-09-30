package main

import "fmt"

func SafeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("ошибка при делении: %v", r)
		}
	}()

	result = a / b
	return result, nil
}

func main() {
	pairs := [][2]int{
		{10, 2},
		{9, 0},
		{20, 4},
	}

	for _, p := range pairs {
		result, err := SafeDivide(p[0], p[1])
		if err != nil {
			fmt.Printf("%d / %d: %v\n", p[0], p[1], err)
			continue
		}
		fmt.Printf("%d / %d = %d\n", p[0], p[1], result)
	}

	fmt.Println("программа не упала")
}
