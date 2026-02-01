package http

import (
	"net/http"

	"distributed-url-shortener/internal/service"
)

func RedirectHandler(svc Shortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Path[1:]

		url, err := svc.Resolve(code)

		if err == service.ErrNotFound {
			http.NotFound(w, r)
			return
		}

		if err == service.ErrServiceUnavailable {
			http.Error(w, "Service temporarily unavailable", http.StatusServiceUnavailable)
			return
		}

		http.Redirect(w, r, url, http.StatusFound)
	}
}
