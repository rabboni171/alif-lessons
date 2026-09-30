package tracker

import (
	"time"

	"github.com/google/uuid"
)

// StartedAt - время запуска приложения. Выставляется один раз в init(),
// поэтому у всех "сессий", созданных за время работы программы, оно
// будет одинаковым.
var StartedAt time.Time

func init() {
	StartedAt = time.Now()
}

// NewSessionID возвращает новый уникальный ID сессии.
func NewSessionID() string {
	return uuid.NewString()
}
