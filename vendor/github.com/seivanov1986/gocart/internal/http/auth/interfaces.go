package auth

import (
	"net/http"
)

const (
	authLoginBodySize = 1024
)

type Handle interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Ping(w http.ResponseWriter, r *http.Request)
}
