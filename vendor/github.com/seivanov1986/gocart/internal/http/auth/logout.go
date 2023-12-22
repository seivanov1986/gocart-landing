package auth

import (
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
)

func (h *handle) Logout(w http.ResponseWriter, r *http.Request) {
	tokenSource := r.Context().Value("authorization")
	token, exists := tokenSource.(string)

	if !exists {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	err := h.service.Logout(token)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK)
}
