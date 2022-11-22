package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"chatty/chatty/app/config"
)

func Connection(ctx context.Context, pgCfg config.PG) (*pgxpool.Pool, error) {
	cfgStr := fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s port=%d pool_max_conns=%d",
		pgCfg.User, pgCfg.Password, pgCfg.Host, pgCfg.DbName, pgCfg.Port, pgCfg.PoolMax)

	cfg, err := pgxpool.ParseConfig(cfgStr)
	if err != nil {
		return nil, err
	}

	pgConn, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return pgConn, nil
}
