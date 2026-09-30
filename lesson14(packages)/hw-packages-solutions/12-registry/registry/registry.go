package registry

// Handler - общий интерфейс для обработчиков текста.
type Handler interface {
	Handle(text string) string
}

var handlers = map[string]Handler{}

// Register добавляет обработчик в общий реестр под именем name.
func Register(name string, h Handler) {
	handlers[name] = h
}

// All возвращает карту зарегистрированных обработчиков.
func All() map[string]Handler {
	return handlers
}
