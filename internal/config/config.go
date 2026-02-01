package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port       string
	RedisAddr  string
	CacheSize  int
	RateLimit  int
	RateWindow int
}

func Load() *Config {
	return &Config{
		Port:       getEnv("PORT", "8080"),
		RedisAddr:  getEnv("REDIS_ADDR", "127.0.0.1:6379"),
		CacheSize:  getEnvInt("CACHE_SIZE", 100000),
		RateLimit:  getEnvInt("RATE_LIMIT", 10),
		RateWindow: getEnvInt("RATE_WINDOW", 60),
	}
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
