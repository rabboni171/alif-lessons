package main

import "fmt"

func Classify(v any) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("целое число: %d", x)
	case float64:
		return fmt.Sprintf("дробное число: %v", x)
	case string:
		return fmt.Sprintf("строка длиной %d символов", len(x))
	case bool:
		return "логическое значение"
	case nil:
		return "пусто"
	default:
		return fmt.Sprintf("неизвестный тип: %T", x)
	}
}

func main() {
	values := []any{42, 3.14, "привет", true, nil, []int{1, 2, 3}}

	for _, v := range values {
		fmt.Println(Classify(v))
	}
}
