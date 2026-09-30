package main

import "fmt"

type Logger struct {
	prefix string
	count  int
}

func NewLogger(prefix string) *Logger {
	return &Logger{prefix: prefix}
}

func (l *Logger) Log(msg string) {
	l.count++
	fmt.Printf("[%s] %s (#%d)\n", l.prefix, msg, l.count)
}

func main() {
	appLogger := NewLogger("APP")
	dbLogger := NewLogger("DB")

	appLogger.Log("сервер запущен")
	dbLogger.Log("подключение установлено")
	appLogger.Log("получен запрос")
	dbLogger.Log("запрос выполнен")
}
