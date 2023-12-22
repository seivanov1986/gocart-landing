package common

import (
	"context"
)

type Service interface {
	Process(ctx context.Context, path string) ([]byte, error)
}
