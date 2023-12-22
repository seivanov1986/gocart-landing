package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	user2 "github.com/seivanov1986/gocart/internal/service/user"
)

type UserUpdateRpcIn struct {
	ID       int64  `json:"id"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}

type UserUpdateRpcOut struct {
	ID *int64
}

type UserUpdateError struct {
	Error string
}

func (u *handle) Update(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	CreateListInput, err := validateUserUpdate(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	err = u.service.Update(r.Context(), *CreateListInput)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError, UserCreateError{
			Error: err.Error(),
		})
		return
	}

	helpers.HttpResponse(w, http.StatusOK)
}

func validateUserUpdate(bodyBytes []byte) (*user2.UserUpdateInput, error) {
	var (
		out             user2.UserUpdateInput
		userUpdateRpcIn UserUpdateRpcIn
	)

	if err := json.Unmarshal(bodyBytes, &userUpdateRpcIn); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	out.Password = userUpdateRpcIn.Password
	out.ID = userUpdateRpcIn.ID

	return &out, nil
}
