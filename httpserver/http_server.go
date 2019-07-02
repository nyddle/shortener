package httpserver

import (
	"fmt"
	"net/http"

	"shortener/service"

	"github.com/gorilla/mux"
)

// HttpServer represents transport layer of out app
type HttpServer struct {
	httpPort int
	router   http.Handler
	s        *service.ShortenerService
}

// Start fires up the http server
func (s *HttpServer) Start() error {
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", s.httpPort), s.router)
}

// NewHTTPServer returns http server that wraps shortener business logic
func NewHTTPServer(ss *service.ShortenerService, port int) *HttpServer {

	r := mux.NewRouter()
	hs := HttpServer{router: r, httpPort: port, s: ss}

	r.HandleFunc("/{shortened}", ss.ResolverHandle)
	http.Handle("/", r)

	return &hs
}
