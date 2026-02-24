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

	// =============================
	// CONFIG
	// =============================

	port := getEnv("PORT", "8080")
	cacheSize := getEnvInt("CACHE_SIZE", 100000)

	// Rate limiting config
	enableRateLimit := getEnv("ENABLE_RATE_LIMIT", "true") == "true"
	rateLimit := getEnvInt("RATE_LIMIT", 1000)
	rateWindowSeconds := getEnvInt("RATE_WINDOW_SECONDS", 60)

	// =============================
	// METRICS
	// =============================

	metrics.Register()

	// =============================
	// REDIS
	// =============================

	redisClient := config.NewRedisClient()

	// =============================
	// CACHE
	// =============================

	lruCache, err := cache.NewLRUCache(cacheSize)
	if err != nil {
		log.Fatal("❌ Failed to initialize cache:", err)
	}

	// =============================
	// SERVICE
	// =============================

	shortener := service.NewShortenerService(redisClient, lruCache)

	// =============================
	// RATE LIMITER (Optional)
	// =============================

	var handler http.Handler

	if enableRateLimit {
		log.Println("🛡 Rate limiter ENABLED")
		rateLimiter := middleware.NewRateLimiter(
			redisClient,
			rateLimit,
			time.Duration(rateWindowSeconds)*time.Second,
		)
		handler = httpserver.NewRouter(shortener, rateLimiter)
	} else {
		log.Println("⚠️ Rate limiter DISABLED (dev mode)")
		handler = httpserver.NewRouter(shortener, nil)
	}

	// =============================
	// SERVER
	// =============================

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Println("✅ HTTP server listening on :" + port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("❌ Server failed:", err)
	}
}

// =============================
// Helpers
// =============================

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
