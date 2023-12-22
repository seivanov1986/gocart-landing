package auth

import (
	"net/http"
)

type Middleware interface {
	Handle(next http.Handler) http.Handler
}

type SessionManager interface {
	Exists(keys ...string) (bool, error)
}
