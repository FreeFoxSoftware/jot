package store

func List(isGlobal bool, all bool) ([]Note, error) {
	s, err := loadAll()

	if err != nil {
		return nil, err
	}

	if all {
		return allNotes(s), nil
	}
	key, err := storeKey(isGlobal)

	if err != nil {
		return nil, err
	}

	return s.Notes[key], nil
}
