package main

import "fmt"

var AppName string

// init выполняется автоматически до main() - AppName уже готова
// к моменту запуска main, вручную ничего присваивать не нужно.
func init() {
	AppName = "MyApp v1.0"
}

func main() {
	fmt.Println(AppName)
}
