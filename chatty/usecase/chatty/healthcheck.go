package chatty

import (
	"context"
)

func (e *chatUseCase) HealthCheck(ctx context.Context) error {
	return e.userRepo.HealthCheck(ctx)
}
