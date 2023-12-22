package category

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/repository/category"
	category2 "github.com/seivanov1986/gocart/internal/service/category"
)

type CategoryReadRpcIn struct {
	ID int64 `json:"id"`
}

type CategoryReadRpcOut struct {
	Row category.CategoryReadRow
}

func (u *handle) Read(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	UserRow, err := validateCategoryRead(bodyBytes)
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

	helpers.HttpResponse(w, http.StatusOK, CategoryReadRpcOut{
		Row: *listOut,
	})
}

func validateCategoryRead(bodyBytes []byte) (*category2.CategoryReadIn, error) {
	listInt := category2.CategoryReadIn{}
	userReadRpcIn := CategoryReadRpcIn{}

	err := json.Unmarshal(bodyBytes, &userReadRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.ID = userReadRpcIn.ID

	return &listInt, nil
}
