package product

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/repository/product"
	product2 "github.com/seivanov1986/gocart/internal/service/product"
)

type ProductReadRpcIn struct {
	ID int64 `json:"id"`
}

type ProductReadRpcOut struct {
	Row product.ProductReadRow
}

func (u *handle) Read(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	UserRow, err := validateProductRead(bodyBytes)
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

	helpers.HttpResponse(w, http.StatusOK, ProductReadRpcOut{
		Row: *listOut,
	})
}

func validateProductRead(bodyBytes []byte) (*product2.ProductReadIn, error) {
	listInt := product2.ProductReadIn{}
	userReadRpcIn := ProductReadRpcIn{}

	err := json.Unmarshal(bodyBytes, &userReadRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.ID = userReadRpcIn.ID

	return &listInt, nil
}
