package delivery

import "github.com/pkg/errors"

var ErrBadContentType = errors.New("content-Type application/json is missing")
var ErrBadRequestBody = errors.New("request body is incorrect")
var ErrUnauthorizied = errors.New("authorization is required")
var ErrBadRequest = errors.New("bad request")
