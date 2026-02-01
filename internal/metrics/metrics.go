package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	HttpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total HTTP requests",
		},
		[]string{"path", "method", "status"},
	)

	HttpLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latency",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)

	RedisErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "redis_errors_total",
			Help: "Total Redis errors",
		},
	)
)

func Register() {
	prometheus.MustRegister(
		HttpRequests,
		HttpLatency,
		RedisErrors,
	)
}
