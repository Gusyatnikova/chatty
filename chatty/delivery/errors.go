package delivery

import "github.com/pkg/errors"

var ErrBadContentType = errors.New("Content-Type application/json is missing")
var ErrBadRequestBody = errors.New("Request body is incorrect")
var ErrUnauthorizied = errors.New("Authorization is required")
