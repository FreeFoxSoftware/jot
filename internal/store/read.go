package store

import (
	"fmt"
)

func Read(id int) (Note, error) {
	s, err := loadAll()
	if err != nil {
		return Note{}, err
	}

	for _, notes := range s.Notes {
		for _, note := range notes {
			if note.ID == id {
				return note, nil
			}
		}
	}

	return Note{}, fmt.Errorf("note with ID %d not found", id)
}
