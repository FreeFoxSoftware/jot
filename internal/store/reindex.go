package store

import "sort"

func Reindex() (int, error) {
	s, err := loadAll()
	if err != nil {
		return 0, err
	}

	type entry struct {
		key  string
		idx  int
		note Note
	}

	var entries []entry
	for key, notes := range s.Notes {
		for i, note := range notes {
			entries = append(entries, entry{key, i, note})
		}
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].note.ID < entries[j].note.ID
	})

	for i := range entries {
		entries[i].note.ID = i + 1
	}

	newNotes := make(map[string][]Note, len(s.Notes))
	for _, e := range entries {
		newNotes[e.key] = append(newNotes[e.key], e.note)
	}

	s.Notes = newNotes
	s.NextID = len(entries)

	return len(entries), saveAll(s)
}
