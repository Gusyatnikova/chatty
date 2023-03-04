package usecase

import (
	"github.com/pkg/errors"

	"chatty/chatty/repository"
)

var ErrInternalError = errors.New("internal server error")

var (
	ErrNoUser           = errors.New("user not found")
	ErrUserDuplication  = errors.New("such user is already exists")
	ErrDataDuplication  = errors.New("email and phone must be unique")
	ErrUserUnauthorized = errors.New("authorization error")
)

func ErrorHandling(err error) error {
	if errors.Is(err, repository.ErrNotFound) {
		return ErrNoUser
	} else if errors.Is(err, repository.ErrDuplicate) {
		return ErrDataDuplication
	}

	return err
}
