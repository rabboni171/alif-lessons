package main

import "fmt"

func AsString(v any) (string, bool) {
	s, ok := v.(string)
	return s, ok
}

func main() {
	values := []any{"привет", 42, 3.14, "мир", true}

	for _, v := range values {
		s, ok := AsString(v)
		fmt.Println(s, ok)
	}
}
