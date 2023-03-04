package repository

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

var (
	ErrNotFound  = errors.New("requested data is not found")
	ErrDuplicate = errors.New("duplicate unique data")
)

func ErrorHandling(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return ErrDuplicate
		}
	}

	return err
}
