package main

import (
	"fmt"
	"time"

	"session/tracker"
)

func main() {
	sessions := []string{
		tracker.NewSessionID(),
		tracker.NewSessionID(),
		tracker.NewSessionID(),
	}

	for _, id := range sessions {
		fmt.Printf("сессия %s, запущена в %s\n", id, tracker.StartedAt.Format(time.Kitchen))
	}
}
