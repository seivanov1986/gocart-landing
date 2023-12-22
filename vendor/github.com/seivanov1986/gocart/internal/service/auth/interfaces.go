package user

import (
	"context"
	"time"
)

type Service interface {
	Login(ctx context.Context, in AuthLoginIn) (*string, error)
	Logout(sessionId string) error
}

type SessionManager interface {
	Del(keys ...string) (bool, error)
	Set(key string, value interface{}, expiration time.Duration) error
}
