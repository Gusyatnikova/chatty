package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"chatty/chatty/entity"
	"chatty/chatty/repository"
)

const (
	GetUserByLoginSqlCmd = `select id, login, password, email, phone_number from public.user
							where login = $1`
	GetUserByIDSqlCmd = `select id, login, password, email, phone_number from public.user
							where id = $1`
)

func (e *pgUserRepo) GetUserByLogin(ctx context.Context, login entity.UserLogin) (user entity.User, err error) {
	row := e.db.QueryRow(ctx, GetUserByLoginSqlCmd, login)
	user, err = e.scanUser(row)
	if err != nil {
		return user, repository.ErrorHandling(err)
	}

	return user, errors.Wrap(err, "Err in: pgUserRepo.GetUserByLogin.ScanUser(): ")
}

func (e *pgUserRepo) GetUserByID(ctx context.Context, id entity.UserID) (user entity.User, err error) {
	row := e.db.QueryRow(ctx, GetUserByIDSqlCmd, id)
	user, err = e.scanUser(row)
	if err != nil {
		return user, repository.ErrorHandling(err)
	}

	return user, errors.Wrap(err, "Err in: pgUserRepo.GetUserByID.ScanUser(): ")
}

func (e *pgUserRepo) scanUser(row pgx.Row) (entity.User, error) {
	var (
		user  = entity.User{}
		idStr = ""
	)

	err := row.Scan(&idStr, &user.Creds.Login, &user.Creds.Password,
		&user.Contacts.Email, &user.Contacts.PhoneNumber)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user, repository.ErrNotFound
		}
		return user, errors.Wrap(err, "Err in: pgUserRepo.ScanUser.Scan()")
	}

	user.SetID(entity.UserID(idStr))

	return user, nil
}
