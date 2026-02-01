package middleware

import (
	"net/http"
	"time"

	"distributed-url-shortener/internal/metrics"
)

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rec := &statusRecorder{ResponseWriter: w, status: 200}
		next.ServeHTTP(rec, r)

		elapsed := time.Since(start).Seconds()

		metrics.HttpRequests.WithLabelValues(
			r.URL.Path,
			r.Method,
			http.StatusText(rec.status),
		).Inc()

		metrics.HttpLatency.WithLabelValues(r.URL.Path).Observe(elapsed)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}
