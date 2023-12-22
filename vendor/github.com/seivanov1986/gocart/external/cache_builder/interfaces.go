package cache_builder

import (
	"context"
)

type Service interface {
	Make(ctx context.Context) error
}
