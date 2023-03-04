package jwt

import "github.com/pkg/errors"

var (
	ErrUnexpectedSigningMethod = errors.New("Unexpected signing method")
)
