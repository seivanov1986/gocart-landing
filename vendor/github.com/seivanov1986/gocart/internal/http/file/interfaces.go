package file

import (
	"net/http"
)

type Handle interface {
	Static(w http.ResponseWriter, r *http.Request)
	AdminStatic(w http.ResponseWriter, r *http.Request)
	Dynamic(w http.ResponseWriter, r *http.Request)
}
