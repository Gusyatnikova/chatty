package repository

import "github.com/pkg/errors"

var (
	ErrNotFound  = errors.New("Requested data is not found")
	ErrDuplicate = errors.New("Duplicate unique data")
)
