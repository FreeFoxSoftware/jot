package store

import (
	"fmt"
)

func Delete(id int) error {
	s, err := loadAll()
	if err != nil {
		return err
	}

	for key, notes := range s.Notes {
		for i, note := range notes {
			if note.ID == id {
				s.Notes[key] = append(notes[:i], notes[i+1:]...)
				return saveAll(s)
			}
		}
	}

	return fmt.Errorf("note with ID %d not found", id)
}
