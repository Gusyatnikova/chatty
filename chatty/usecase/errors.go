package usecase

import "github.com/pkg/errors"

var ErrInternalError = errors.New("Internal server error")

var (
	ErrNoUser          = errors.New("User not found")
	ErrUserDuplication = errors.New("Such user is already exists")
	ErrDataDuplication = errors.New("Email and phone must be unique")
)

var ErrUserUnauthorized = errors.New("Authorization Error")
