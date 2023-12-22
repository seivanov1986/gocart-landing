package common

import (
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
)

func (u *handle) Process(w http.ResponseWriter, r *http.Request) {
	bytes, err := u.service.Process(r.Context(), r.URL.Path)
	if err != nil {
		helpers.HttpResponse(w, http.StatusNotFound)
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}
}
