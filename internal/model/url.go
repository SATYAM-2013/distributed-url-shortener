package model

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Code string `json:"code"`
}
