package product_to_category

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
)

type ProductToCategoryListRpcIn struct {
}

type ProductToCategoryListError struct {
}

func (u *handle) List(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	_, err = validateProductToCategoryList(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	helpers.HttpResponse(w, http.StatusOK)
}

func validateProductToCategoryList(bodyBytes []byte) ([]int64, error) {
	listInt := []int64{}
	userListRpcIn := ProductToCategoryListRpcIn{}

	err := json.Unmarshal(bodyBytes, &userListRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return listInt, nil
}
