package service

import (
	"testing"

	"distributed-url-shortener/internal/config"
)

func setupTestService(t *testing.T) ShortenerService {
	rdb := config.NewRedisClient()

	// cleanup before each test
	rdb.FlushDB(config.Ctx)

	return NewShortenerService(rdb)
}
func TestShorten_GeneratesCode(t *testing.T) {
	svc := setupTestService(t)

	url := "https://google.com"
	code := svc.Shorten(url)

	if code == "" {
		t.Fatal("expected non-empty short code")
	}
}
func TestResolve_ReturnsOriginalURL(t *testing.T) {
	svc := setupTestService(t)

	url := "https://google.com"
	code := svc.Shorten(url)

	resolvedURL, ok := svc.Resolve(code)

	if !ok {
		t.Fatal("expected code to resolve")
	}

	if resolvedURL != url {
		t.Fatalf("expected %s, got %s", url, resolvedURL)
	}
}
func TestResolve_InvalidCode(t *testing.T) {
	svc := setupTestService(t)

	_, ok := svc.Resolve("invalid123")

	if ok {
		t.Fatal("expected resolve to fail for invalid code")
	}
}
