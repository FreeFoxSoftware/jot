package store

import "time"

type Note struct {
	ID        int
	Title     string
	Body      string
	Global    bool
	CreatedAt time.Time
}
