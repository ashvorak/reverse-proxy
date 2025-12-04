package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"reverse-proxy/internal/router"
)

type Server struct {
	router *router.Router
}

func NewServer(router *router.Router) (*Server, error) {
	if router == nil {
		return nil, errors.New("can't create server without router")
	}

	server := &Server{
		router: router,
	}

	return server, nil
}

func (s *Server) Start() {
	log.Fatal(http.ListenAndServe(":8080", s))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upstream, ok := s.router.Match(r.URL.Path)
	if !ok {
		fmt.Fprint(w, "No valid upstream.")
		return
	}

	fmt.Fprintf(w, "Upstream for handling: %s", upstream)
}
