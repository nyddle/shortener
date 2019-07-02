package myshortener

import (
	"encoding/base64"
	"sync"
)

func shorten(url string) string {
	return base64.StdEncoding.EncodeToString([]byte(url))
}

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
	defer s.RUnlock()
	return s.storage[url]
}

func NewMyShortener() *MyShortener {

	return &MyShortener{storage: make(map[string]string)}
}
