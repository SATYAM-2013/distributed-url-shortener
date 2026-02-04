package http

import (
	"encoding/json"
	"net/http"

	"distributed-url-shortener/internal/http/middleware"
	"distributed-url-shortener/internal/service"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type RootResponse struct {
	Service   string            `json:"service"`
	Status    string            `json:"status"`
	Endpoints map[string]string `json:"endpoints"`
}

func NewRouter(
	svc *service.ShortenerService,
	rl *middleware.RateLimiter,
) http.Handler {

	mux := http.NewServeMux()

	// ----------------------
	// Root (API info)
	// ----------------------
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(RootResponse{
			Service: "distributed-url-shortener",
			Status:  "running",
			Endpoints: map[string]string{
				"POST /shorten": "Create short URL",
				"GET /{code}":   "Redirect to original URL",
				"GET /health":   "Health check",
				"GET /metrics":  "Prometheus metrics",
			},
		})
	})

	// ----------------------
	// Health
	// ----------------------
	mux.HandleFunc("/health", HealthHandler)

	// ----------------------
	// Metrics
	// ----------------------
	mux.Handle("/metrics", promhttp.Handler())

	// ----------------------
	// Shorten URL
	// ----------------------
	mux.Handle(
		"/shorten",
		middleware.Metrics(
			rl.Middleware(
				ShortenHandler(svc),
			),
		),
	)

	// ----------------------
	// Redirect short URL
	// ----------------------
	mux.Handle(
		"/",
		middleware.Metrics(
			RedirectHandler(svc),
		),
	)

	return mux
}
