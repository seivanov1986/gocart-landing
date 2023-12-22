package observer

import (
	"context"
)

const (
	serviceBasePath = "service_base_path"
)

func GetServiceBasePath(ctx context.Context) string {
	result := ""

	if value, ok := ctx.Value(serviceBasePath).(string); ok {
		result = value
	}

	return result
}

func SetServiceBasePath(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, serviceBasePath, path)
}
