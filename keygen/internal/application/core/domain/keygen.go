package domain

import "time"

type KeyGenLogEntry struct {
	ShortUrl  string
	CreatedAt time.Time
}

// NewKeygenLogEntry creates a new KeyGenLogEntry
func NewKeygenLogEntry(shortUrl string) KeyGenLogEntry {
	return KeyGenLogEntry{
		ShortUrl:  shortUrl,
		CreatedAt: time.Now(),
	}
}
