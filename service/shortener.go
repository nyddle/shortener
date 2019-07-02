package service

import (
	"bytes"
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
	address string
}

// NewShortenerService creates a new shortener service
func NewShortenerService(shortener Shortener, address string) *ShortenerService {
	return &ShortenerService{shortener, address}
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
		shortened := ss.address + "/" + ss.Shortener.Shorten(buf.String())
		w.Write([]byte(shortened))
	}
}
