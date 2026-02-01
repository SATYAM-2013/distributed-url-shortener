package http

import (
	"log"
	"net/http"
)

type Server struct {
	addr   string
	router http.Handler
}

func NewServer(addr string, router http.Handler) *Server {
	return &Server{
		addr:   addr,
		router: router,
	}
}

func (s *Server) Start() error {
	log.Println("HTTP server listening on", s.addr)
	return http.ListenAndServe(s.addr, s.router)
}
