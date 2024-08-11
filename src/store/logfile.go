package store

import (
	"log"
	"os"
)

type LogFileStore struct {
	latestContent string
	file          *os.File
	logger        *log.Logger
}

func NewLogFileStore(filename string) *LogFileStore {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "", log.LstdFlags)

	store := &LogFileStore{
		file:   file,
		logger: logger,
	}
	return store
}

func (s *LogFileStore) GetLatestContent() string {
	return ""
}

func (s *LogFileStore) SetLatestContent(content string) {
	if content != s.latestContent {
		s.logger.Printf("%s", content)
		s.latestContent = content
	}
}

func (s *LogFileStore) Close() {
    s.file.Close()
}
