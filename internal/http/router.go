package http

import (
	"net/http"

	"distributed-url-shortener/internal/service"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	shortenerSvc := service.NewShortenerService()

	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/shorten", ShortenHandler(shortenerSvc))

	// ⚠️ MUST BE LAST
	mux.HandleFunc("/", RedirectHandler(shortenerSvc))

	return mux
}
