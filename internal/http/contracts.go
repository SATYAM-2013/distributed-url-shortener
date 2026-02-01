package http

type Shortener interface {
	Shorten(url string) (string, error)
	Resolve(code string) (string, error)
}
