package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	authService "github.com/seivanov1986/gocart/internal/service/auth"
)

type LoginRpcIn struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginRpcOut struct {
	Token *string `json:"token"`
}

func (h *handle) Login(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, authLoginBodySize))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	authLoginInput, err := validateAuthLogin(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(r.Context(), *authLoginInput)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK, LoginRpcOut{
		Token: token,
	})
}

func validateAuthLogin(bodyBytes []byte) (*authService.AuthLoginIn, error) {
	var (
		out            authService.AuthLoginIn
		authLoginRpcIn LoginRpcIn
	)

	if err := json.Unmarshal(bodyBytes, &authLoginRpcIn); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	out.Login = authLoginRpcIn.Login
	out.Password = authLoginRpcIn.Password

	return &out, nil
}
