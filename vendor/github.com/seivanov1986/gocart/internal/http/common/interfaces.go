package common

import (
	"net/http"
)

type Handle interface {
	Process(w http.ResponseWriter, r *http.Request)
}
