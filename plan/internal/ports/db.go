package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/plan/internal/application/domain"
)

type DBPort interface {
	Create(ctx context.Context, plans []domain.Plan) error
	GetAll(ctx context.Context) ([]domain.Plan, error)
	Get(ctx context.Context, id string) (domain.Plan, error)
}
