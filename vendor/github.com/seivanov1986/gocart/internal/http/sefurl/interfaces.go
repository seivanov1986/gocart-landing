package sefurl

import (
	"net/http"
)

type Handle interface {
	List(w http.ResponseWriter, r *http.Request)
}
