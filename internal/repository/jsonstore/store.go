package jsonstore

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

type Store struct {
	root string
	mu   sync.RWMutex
}

func NewStore(root string) *Store { return &Store{root: root} }

func (s *Store) read(name string, v any) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	f, err := os.Open(filepath.Join(s.root, name))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrNotFound
		}
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(v)
}

func (s *Store) write(name string, v any) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	tmp := filepath.Join(s.root, name+".tmp")
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	if err := json.NewEncoder(f).Encode(v); err != nil {
		f.Close()
		return err
	}
	f.Close()
	return os.Rename(tmp, filepath.Join(s.root, name))
}

var ErrNotFound = errors.New("not found")
