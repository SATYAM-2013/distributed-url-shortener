package http

import (
	"net/http"

	"distributed-url-shortener/internal/service"
)

func RedirectHandler(svc service.ShortenerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		code := r.URL.Path[1:] // remove leading "/"
		if code == "" {
			http.NotFound(w, r)
			return
		}

		url, ok := svc.Resolve(code)
		if !ok {
			http.NotFound(w, r)
			return
		}

		http.Redirect(w, r, url, http.StatusFound)
	}
}
