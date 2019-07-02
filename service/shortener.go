package service

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Shortener implements business logic for the shortener service
type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type ShortenerService struct {
	Shortener
}

// NewShortenerService creates a new shortener service
func NewShortenerService(shortener Shortener) *ShortenerService {
	return &ShortenerService{shortener}
}

func (ss *ShortenerService) ResolverHandle(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		if url, ok := vars["shortened"]; ok {
			ss.Shortener.Resolve(string(url))
			http.Redirect(w, r, ss.Shortener.Resolve(string(url)), http.StatusSeeOther)
		} else {
			w.WriteHeader(404)
		}
	case "POST":
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		shortened := fmt.Sprintf("http://127.0.0.1:5555/") + ss.Shortener.Shorten(buf.String())
		w.Write([]byte(shortened))
	}
}

// ShortenHandle returns shortened url in response body
func (ss *ShortenerService) ShortenHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world"))
}
