package delivery

import "github.com/pkg/errors"

var (
	ErrBadContentType = errors.New("Content-Type application/json is missing")
	ErrBadRequestBody = errors.New("Request body is incorrect")
	ErrUnauthorizied  = errors.New("Authorization is required")
)
