package storage

import "errors"

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLAlreadyExist = errors.New("url already exists")
)