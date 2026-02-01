package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"distributed-url-shortener/internal/config"
	"distributed-url-shortener/internal/service"
)

func setupTestServer(t *testing.T) *httptest.Server {
	rdb := config.NewRedisClient()
	rdb.FlushDB(config.Ctx) // clean state

	shortenerSvc := service.NewShortenerService(rdb)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/shorten", ShortenHandler(shortenerSvc))
	mux.HandleFunc("/", RedirectHandler(shortenerSvc))

	return httptest.NewServer(mux)
}
func TestHealthEndpoint(t *testing.T) {
	server := setupTestServer(t)
	defer server.Close()

	resp, err := http.Get(server.URL + "/health")
	if err != nil {
		t.Fatalf("failed to call health endpoint: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", resp.StatusCode)
	}
}
func TestShortenEndpoint(t *testing.T) {
	server := setupTestServer(t)
	defer server.Close()

	body := map[string]string{
		"url": "https://google.com",
	}

	jsonBody, _ := json.Marshal(body)

	resp, err := http.Post(
		server.URL+"/shorten",
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		t.Fatalf("failed to call shorten endpoint: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", resp.StatusCode)
	}

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	if result["code"] == "" {
		t.Fatal("expected non-empty short code")
	}
}
func TestRedirectEndpoint(t *testing.T) {
	server := setupTestServer(t)
	defer server.Close()

	// Step 1: shorten URL
	body := map[string]string{
		"url": "https://google.com",
	}
	jsonBody, _ := json.Marshal(body)

	resp, err := http.Post(
		server.URL+"/shorten",
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		t.Fatalf("failed to shorten url: %v", err)
	}

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	code := result["code"]

	// Step 2: request redirect
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	redirectResp, err := client.Get(server.URL + "/" + code)
	if err != nil {
		t.Fatalf("failed to call redirect: %v", err)
	}

	if redirectResp.StatusCode != http.StatusFound {
		t.Fatalf("expected 302 redirect, got %d", redirectResp.StatusCode)
	}

	location := redirectResp.Header.Get("Location")
	if location != "https://google.com" {
		t.Fatalf("expected redirect to google.com, got %s", location)
	}
}
