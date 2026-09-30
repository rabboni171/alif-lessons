package reverse

import "registrydemo/registry"

// Handler - обработчик, разворачивающий текст.
type Handler struct{}

func (Handler) Handle(text string) string {
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func init() {
	registry.Register("reverse", Handler{})
}
