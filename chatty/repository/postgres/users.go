package postgres

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"chatty/chatty/entity"
	"chatty/chatty/repository"
)

const (
	AddUserSqlCmd        = `insert into public.user values($1, $2, $3, $4, $5, now())`
	GetUserByLoginSqlCmd = `select id, login, password, email, phone_number from public.user
							where login = $1`
)

func (e *pgUserRepo) AddUser(ctx context.Context, user entity.User) error {
	addUserFn := func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, AddUserSqlCmd,
			user.ID, user.Creds.Login, user.Creds.Password,
			user.Contacts.Email, user.Contacts.PhoneNumber)
		return errors.Wrap(err, "Err in pgUserRepo.AddUser.Exec()")
	}

	return errorHandling(e.RunTx(ctx, addUserFn))
}

func (e *pgUserRepo) GetUserByLogin(ctx context.Context, login entity.UserLogin) (entity.User, error) {
	row := e.db.QueryRow(ctx, GetUserByLoginSqlCmd, login)
	user, err := e.scanUser(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, repository.ErrNotFound
		}
	}

	return user, errors.Wrap(err, "Err in: pgUserRepo.GetUserByLogin.ScanUser(): ")
}

func (e *pgUserRepo) scanUser(row pgx.Row) (entity.User, error) {
	user := entity.User{}

	err := row.Scan(&user.ID, &user.Creds.Login, &user.Creds.Password,
		&user.Contacts.Email, &user.Contacts.PhoneNumber)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, repository.ErrNotFound
		}
		return user, errors.Wrap(err, "Err in: pgUserRepo.ScanUser.Scan()")
	}

	return user, nil
}

func errorHandling(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return repository.ErrNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return repository.ErrDuplicate
		}
	}

	return err
}
