package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const globalKey = "global"

type Store struct {
	NextID int               `json:"nextID"`
	Notes  map[string][]Note `json:"notes"`
}

func storeKey(global bool) (string, error) {
	if global {
		return globalKey, nil
	}
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("could not get working directory: %w", err)
	}
	return cwd, nil
}

func loadAll() (Store, error) {
	data, err := os.ReadFile(storePath())
	if os.IsNotExist(err) {
		return Store{
			Notes: map[string][]Note{},
		}, nil
	}
	if err != nil {
		return Store{
			Notes: map[string][]Note{},
		}, fmt.Errorf("could not read store: %w", err)
	}

	store := Store{}
	if err := json.Unmarshal(data, &store); err != nil {
		return Store{
			Notes: map[string][]Note{},
		}, fmt.Errorf("could not parse store: %w", err)
	}
	return store, nil
}

func saveAll(store Store) error {
	path := storePath()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("could not create store directory: %w", err)
	}

	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return fmt.Errorf("could not encode notes: %w", err)
	}

	return os.WriteFile(path, data, 0644)
}

func storePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".jot", "notes.json")
}

func allNotes(s Store) []Note {
	var all []Note
	for _, notes := range s.Notes {
		all = append(all, notes...)
	}
	return all
}
