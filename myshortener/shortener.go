package myshortener

import (
	"encoding/base64"
	"sync"
)

func shorten(url string) string {
	return base64.StdEncoding.EncodeToString([]byte(url))
}

// MyShortener is basically a wrapper around a map
type MyShortener struct {
	sync.RWMutex
	storage map[string]string
}

// Shorten returns shortened version of a url
func (s *MyShortener) Shorten(url string) string {
	s.Lock()
	defer s.Unlock()

	s.storage[shorten(url)] = url
	return shorten(url)
}

// Resolve returns long version of the url given a shortened one
func (s *MyShortener) Resolve(url string) string {
	s.RLock()
	longURL := s.storage[url]
	defer s.RUnlock()
	return longURL
}

// NewMyShortener returns a map wrapper to lookup shortened urls
func NewMyShortener() *MyShortener {

	return &MyShortener{storage: make(map[string]string)}
}
