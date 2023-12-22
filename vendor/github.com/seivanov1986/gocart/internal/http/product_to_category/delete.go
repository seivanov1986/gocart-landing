package product_to_category

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
)

type ProductToCategoryDeleteRpcIn struct {
}

type ProductToCategoryDeleteError struct {
}

func (u *handle) Delete(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	_, err = validateProductToCategoryDelete(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	helpers.HttpResponse(w, http.StatusOK)
}

func validateProductToCategoryDelete(bodyBytes []byte) ([]int64, error) {
	listInt := []int64{}
	userDeleteRpcIn := ProductToCategoryDeleteRpcIn{}

	err := json.Unmarshal(bodyBytes, &userDeleteRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return listInt, nil
}
