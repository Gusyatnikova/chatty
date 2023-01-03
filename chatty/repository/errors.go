package repository

import "github.com/pkg/errors"

var (
	ErrNotFound  = errors.New("requested data is not found")
	ErrDuplicate = errors.New("duplicate unique data")
)
