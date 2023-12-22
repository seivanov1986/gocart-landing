package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/seivanov1986/gocart/helpers"
	userService "github.com/seivanov1986/gocart/internal/service/user"
)

type UserCreateRpcIn struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}

type UserCreateRpcOut struct {
	ID *int64
}

type UserCreateError struct {
	Error string
}

func (u *handle) Create(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	CreateListInput, err := validateUserCreate(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	userid, err := u.service.Create(r.Context(), *CreateListInput)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError, UserCreateError{
			Error: err.Error(),
		})
		return
	}

	helpers.HttpResponse(w, http.StatusOK, UserCreateRpcOut{
		ID: userid,
	})
}

func validateUserCreate(bodyBytes []byte) (*userService.UserCreateIn, error) {
	var (
		out             userService.UserCreateIn
		userCreateRpcIn UserCreateRpcIn
	)

	if err := json.Unmarshal(bodyBytes, &userCreateRpcIn); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	out.Login = userCreateRpcIn.Login
	out.Email = userCreateRpcIn.Email
	out.Password = userCreateRpcIn.Password
	out.Active = userCreateRpcIn.Active
	out.CreatedAt = time.Now()

	return &out, nil
}
