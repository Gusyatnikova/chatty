package usecase

import "github.com/pkg/errors"

var ErrInternalError = errors.New("internal server error")

var (
	ErrNoUser          = errors.New("user not found")
	ErrUserDuplication = errors.New("such user is already exists")
	ErrDataDuplication = errors.New("email and phone must be unique")
)

var ErrUserUnauthorized = errors.New("authorization error")
