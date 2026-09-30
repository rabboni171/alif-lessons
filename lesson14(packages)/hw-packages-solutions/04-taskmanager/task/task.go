package task

// Task - задача в списке дел.
type Task struct {
	Title string
	Done  bool
}

// New - функция-конструктор. Возвращает Task с Done=false по умолчанию,
// чтобы вызывающий не забыл ни одно поле и не собирал структуру вручную.
func New(title string) Task {
	return Task{Title: title, Done: false}
}
