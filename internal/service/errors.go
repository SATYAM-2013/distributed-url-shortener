package service

import "errors"

var (
	ErrNotFound           = errors.New("not found")
	ErrServiceUnavailable = errors.New("service unavailable")
)
