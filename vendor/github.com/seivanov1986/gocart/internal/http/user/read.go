package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	user2 "github.com/seivanov1986/gocart/internal/repository/user"
)

type UserReadRpcIn struct {
	ID int64 `json:"id"`
}

type UserReadRpcOut struct {
	Row user2.UserReadRow
}

func (u *handle) Read(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	UserRow, err := validateUserRead(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	listOut, err := u.service.Read(r.Context(), *UserRow)
	if err != nil {
		fmt.Println(err)
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK, UserReadRpcOut{
		Row: *listOut,
	})
}

func validateUserRead(bodyBytes []byte) (*int64, error) {
	var listInt int64
	userReadRpcIn := UserReadRpcIn{}

	err := json.Unmarshal(bodyBytes, &userReadRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt = userReadRpcIn.ID

	return &listInt, nil
}
