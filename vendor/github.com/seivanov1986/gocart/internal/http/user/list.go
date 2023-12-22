package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	user2 "github.com/seivanov1986/gocart/internal/repository/user"
	user3 "github.com/seivanov1986/gocart/internal/service/user"
)

type UserListRpcIn struct {
	Page int64 `json:"page"`
}

type UserListRpcOut struct {
	List  []user2.UserListRow
	Total int64
}

func (u *handle) List(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	UserListInput, err := validateUserListV2(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	listOut, err := u.service.List(r.Context(), *UserListInput)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK, UserListRpcOut{
		List:  listOut.List,
		Total: listOut.Total,
	})
}

func validateUserListV2(bodyBytes []byte) (*user3.UserListIn, error) {
	listInt := user3.UserListIn{}
	userListRpcIn := UserListRpcIn{}

	err := json.Unmarshal(bodyBytes, &userListRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.Page = userListRpcIn.Page

	return &listInt, nil
}
