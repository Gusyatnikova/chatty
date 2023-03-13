package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"chatty/chatty/entity"
	"chatty/chatty/repository"
)

const (
	AddUserSqlCmd = `insert into public.user values($1, $2, $3, $4, $5, now())`
)

func (e *pgUserRepo) AddUser(ctx context.Context, user entity.User) error {
	addUserFn := func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, AddUserSqlCmd,
			user.GetID(), user.Creds.Login, user.Creds.Password,
			user.Contacts.Email, user.Contacts.PhoneNumber)
		return errors.Wrap(err, "err in pgUserRepo.AddUser.Exec()")
	}

	return repository.ErrorHandling(e.RunTx(ctx, addUserFn))
}
