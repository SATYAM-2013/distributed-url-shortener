package http

import (
	"net/http"

	"distributed-url-shortener/internal/http/middleware"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(
	svc Shortener,
	rl *middleware.RateLimiter,
) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/health", HealthHandler)
	mux.Handle("/metrics", promhttp.Handler())

	mux.Handle(
		"/shorten",
		middleware.Metrics(
			rl.Middleware(http.HandlerFunc(ShortenHandler(svc))),
		),
	)

	mux.Handle(
		"/",
		middleware.Metrics(
			http.HandlerFunc(RedirectHandler(svc)),
		),
	)

	return mux
}
