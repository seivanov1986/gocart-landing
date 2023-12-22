package common

import (
	"net/http"
)

type Middleware interface {
	Wrapper(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request)
	Handle(next http.Handler) http.Handler
}
