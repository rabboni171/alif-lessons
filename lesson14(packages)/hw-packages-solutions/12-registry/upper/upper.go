package upper

import (
	"strings"

	"registrydemo/registry"
)

// Handler - обработчик, переводящий текст в верхний регистр.
type Handler struct{}

func (Handler) Handle(text string) string {
	return strings.ToUpper(text)
}

// init регистрирует Handler в общем реестре ещё до вызова main() -
// достаточно просто заимпортировать пакет upper, вызывать его функции
// явно не нужно.
func init() {
	registry.Register("upper", Handler{})
}
