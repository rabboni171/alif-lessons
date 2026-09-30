package main

import "github.com/fatih/color"

// PrintCheck печатает результат проверки зелёным (успех) или красным
// (провал) цветом прямо в терминале.
func PrintCheck(name string, passed bool) {
	if passed {
		color.Green("OK: %s", name)
	} else {
		color.Red("FAIL: %s", name)
	}
}

func main() {
	PrintCheck("тест соединения с базой", true)
	PrintCheck("тест авторизации", false)
	PrintCheck("тест валидации email", true)
	PrintCheck("тест отправки уведомлений", false)
}
