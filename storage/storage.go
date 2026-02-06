package storage

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

type Vacancy struct {
	ChatID    int64     `json:"chat_id"`
	MessageID int       `json:"message_id"`
	Text      string    `json:"text"`
	FoundAt   time.Time `json:"found_at"`
}

type FileStorage struct {
	filename string
	mu       sync.Mutex
}

func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{
		filename: filename,
	}
}

// Save appends the vacancy to the file in JSON Lines format
func (s *FileStorage) Save(v Vacancy) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	f, err := os.OpenFile(s.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(v); err != nil {
		return err
	}

	return nil
}
