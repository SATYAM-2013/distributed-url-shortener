package http

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(addr string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (s *Server) Start() {
	log.Printf("HTTP server listening on %s\n", s.httpServer.Addr)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
