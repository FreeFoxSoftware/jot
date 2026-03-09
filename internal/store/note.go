package store

import "time"

type Note struct {
	ID          int
	Title       string
	Description string
	Global      bool
	CreatedAt   time.Time
}
