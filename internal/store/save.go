package store

import (
	"time"
)

func Save(title, body string, global bool) (int, error) {
	s, err := loadAll()
	if err != nil {
		return 0, err
	}

	key, err := storeKey(global)
	if err != nil {
		return 0, err
	}

	s.NextID++
	note := Note{
		ID:        s.NextID,
		Title:     title,
		Body:      body,
		CreatedAt: time.Now(),
	}

	s.Notes[key] = append(s.Notes[key], note)

	return note.ID, saveAll(s)
}
