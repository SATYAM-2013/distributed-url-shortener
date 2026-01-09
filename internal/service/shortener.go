package service

import (
	"crypto/rand"
	"math/big"
	"sync"
)

type ShortenerService interface {
	Shorten(url string) string
	Resolve(code string) (string, bool)
}

type shortenerService struct {
	store map[string]string
	mu    sync.RWMutex
}

func NewShortenerService() ShortenerService {
	return &shortenerService{
		store: make(map[string]string),
	}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const codeLength = 6

func (s *shortenerService) Shorten(url string) string {
	code := generateCode()

	s.mu.Lock()
	s.store[code] = url
	s.mu.Unlock()

	return code
}

func (s *shortenerService) Resolve(code string) (string, bool) {
	s.mu.RLock()
	url, ok := s.store[code]
	s.mu.RUnlock()

	return url, ok
}

func generateCode() string {
	code := make([]byte, codeLength)

	for i := range code {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		code[i] = charset[n.Int64()]
	}

	return string(code)
}
