package main

import (
	"errors"
	"fmt"
)

func ProcessBatch(items []string, process func(string) error) (successCount int, errs []error) {
	for _, item := range items {
		if err := process(item); err != nil {
			errs = append(errs, err)
			continue
		}
		successCount++
	}
	return successCount, errs
}

func main() {
	items := []string{"hello", "", "world", "this is a very long string", "ok"}

	process := func(s string) error {
		if s == "" {
			return errors.New("пустая строка")
		}
		if len(s) > 10 {
			return fmt.Errorf("строка %q слишком длинная", s)
		}
		return nil
	}

	successCount, errs := ProcessBatch(items, process)

	fmt.Println("успешных обработок:", successCount)
	fmt.Println("ошибок:", len(errs))
	for _, err := range errs {
		fmt.Println("  -", err)
	}
}
