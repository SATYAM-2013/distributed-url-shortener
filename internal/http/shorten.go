package http

import (
	"encoding/json"
	"net/http"

	"distributed-url-shortener/internal/service"
)

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	Code string `json:"code"`
}

func ShortenHandler(svc service.ShortenerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()

		var req shortenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		if req.URL == "" {
			http.Error(w, "url is required", http.StatusBadRequest)
			return
		}

		code := svc.Shorten(req.URL)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode(shortenResponse{
			Code: code,
		})
	}
}
