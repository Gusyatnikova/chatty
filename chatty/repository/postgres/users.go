package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"

	"chatty/chatty/entity"
	"chatty/chatty/usecase"
)

type pgUserRepo struct {
	db *pgxpool.Pool
}

func NewPgUserRepo(db *pgxpool.Pool) usecase.ChatUserRepo {
	return &pgUserRepo{db: db}
}

func (e *pgUserRepo) AddUser(ctx context.Context, user entity.User) error { return nil }

func (e *pgUserRepo) CheckUsername(ctx context.Context, username string) (bool, error) {
	return false, nil
}

func (e *pgUserRepo) GetPassword(ctx context.Context, username string) (string, error) {
	return "", nil
}
