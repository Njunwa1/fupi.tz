package ports

import "context"

type KeyGenPort interface {
	GenerateShortUrlKey(ctx context.Context) (string, error)
}
