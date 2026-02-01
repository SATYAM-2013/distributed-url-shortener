package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"distributed-url-shortener/internal/cache"
	"distributed-url-shortener/internal/config"
	httpserver "distributed-url-shortener/internal/http"
	"distributed-url-shortener/internal/http/middleware"
	"distributed-url-shortener/internal/metrics"
	"distributed-url-shortener/internal/service"
)

func main() {
	log.Println("🚀 Distributed URL Shortener starting...")

	// ===== CONFIG =====
	port := getEnv("PORT", "8080")
	cacheSize := getEnvInt("CACHE_SIZE", 100_000)
	redisAddr := getEnv("REDIS_ADDR", "127.0.0.1:6379")

	// ===== METRICS =====
	metrics.Register()

	// ===== REDIS =====
	redisClient := config.NewRedisClient(redisAddr)

	// ===== CACHE =====
	lruCache, err := cache.NewLRUCache(cacheSize)
	if err != nil {
		log.Fatal(err)
	}

	// ===== SERVICE =====
	shortener := service.NewShortenerService(redisClient, lruCache)

	// ===== RATE LIMITER =====
	rateLimiter := middleware.NewRateLimiter(
		redisClient,
		100,
		time.Minute,
	)

	// ===== ROUTER =====
	handler := httpserver.NewRouter(shortener, rateLimiter)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	log.Println("✅ HTTP server listening on :" + port)
	log.Fatal(server.ListenAndServe())
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}
