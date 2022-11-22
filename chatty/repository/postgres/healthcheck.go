package postgres

import (
	"context"
)

func (e *pgUserRepo) HealthCheck(ctx context.Context) error {
	return e.db.Ping(ctx)
}
