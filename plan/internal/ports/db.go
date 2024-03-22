package ports

import "context"

type DBPort interface {
	Create(ctx context.Context)
	Get(ctx context.Context)
}
