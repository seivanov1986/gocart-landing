package auth

import (
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
)

func (u *handle) Ping(w http.ResponseWriter, r *http.Request) {
	helpers.HttpResponse(w, http.StatusOK)
}
