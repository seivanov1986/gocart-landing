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

type ProductListRpcIn struct {
	Page int64 `json:"page"`
}

type ProductListRpcOut struct {
	List  []product.ProductListRow
	Total int64
}

func (u *handle) List(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	PageListInput, err := validateProductList(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	listOut, err := u.service.List(r.Context(), *PageListInput)
	if err != nil {
		fmt.Println(err)
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK, ProductListRpcOut{
		List:  listOut.List,
		Total: listOut.Total,
	})
}

func validateProductList(bodyBytes []byte) (*product2.ProductListIn, error) {
	listInt := product2.ProductListIn{}
	imageListRpcIn := ProductListRpcIn{}

	err := json.Unmarshal(bodyBytes, &imageListRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.Page = imageListRpcIn.Page

	return &listInt, nil
}
