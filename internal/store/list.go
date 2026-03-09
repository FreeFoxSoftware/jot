package store

func List(isGlobal bool) ([]Note, error) {
	s, err := loadAll()
	if err != nil {
		return nil, err
	}

	key, err := storeKey(isGlobal)

	if err != nil {
		return nil, err
	}

	return s.Notes[key], nil
}
