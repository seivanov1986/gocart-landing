package common

import (
	"context"
	"net/http"
)

func (m middleware) Wrapper(
	next func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {
	return m.logic(next)
}

func (m middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(m.logic(next.ServeHTTP))
}

func (m middleware) logic(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
			}
		}()

		ctx := context.WithValue(r.Context(), "service_base_path", m.ServiceBasePath)
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		r = r.WithContext(ctx)
		next(w, r)
	}
}
