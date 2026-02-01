package http

import (
	"encoding/json"
	"net/http"

	"distributed-url-shortener/internal/service"
)

type shortenRequest struct {
	URL string `json:"url"`
}

func ShortenHandler(svc Shortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req shortenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		code, err := svc.Shorten(req.URL)

		if err == service.ErrServiceUnavailable {
			http.Error(w, "Try again later", http.StatusServiceUnavailable)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := map[string]string{
			"short_url": "http://localhost:8080/" + code,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
