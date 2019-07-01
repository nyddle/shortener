package httpserver

import (
	"bytes"
	"fmt"
	"net/http"
	"shortener/service"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	httpPort int
	router   http.Handler
	s        *service.ShortenerService
}

func (hs *HttpServer) resolverHandle(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		if url, ok := vars["shortened"]; ok {
			//w.WriteHeader(302)
			hs.s.Shortener.Resolve(string(url))
			http.Redirect(w, r, hs.s.Shortener.Resolve(string(url)), http.StatusSeeOther)
		} else {
			w.WriteHeader(404)
		}
	case "POST":
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		shortened := fmt.Sprintf("http://127.0.0.1:%d/", hs.httpPort) + hs.s.Shortener.Shorten(buf.String())
		w.Write([]byte(shortened))
	}
}

func (hs *HttpServer) shortenHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world"))
}

func (s *HttpServer) Start() error {
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", s.httpPort), s.router)
}

// NewHTTPServer returns http server that wraps shortener business logic
func NewHTTPServer(sh *service.ShortenerService, port int) *HttpServer {

	r := mux.NewRouter()
	hs := HttpServer{router: r, httpPort: port, s: sh}

	r.HandleFunc("/{shortened}", hs.resolverHandle)
	r.HandleFunc("/{shortened}", hs.shortenHandle)
	http.Handle("/", r)

	return &hs
}
