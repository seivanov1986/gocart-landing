package product

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
)

type ProductDeleteRpcIn struct {
	IDs []int64 `json:"ids"`
}

type ProductDeleteError struct {
	Error string
}

func (u *handle) Delete(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	CreateListInput, err := validateProductDelete(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	err = u.service.Delete(r.Context(), CreateListInput)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError, ProductDeleteError{
			Error: err.Error(),
		})
		return
	}

	helpers.HttpResponse(w, http.StatusOK)
}

func validateProductDelete(bodyBytes []byte) ([]int64, error) {
	listInt := []int64{}
	userCreateRpcIn := ProductDeleteRpcIn{}

	err := json.Unmarshal(bodyBytes, &userCreateRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt = userCreateRpcIn.IDs

	return listInt, nil
}
