package image_to_product

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
)

type ImageToProductDeleteRpcIn struct {
}

type ImageToProductDeleteError struct {
}

func (u *handle) Delete(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	_, err = validateImageToProductDelete(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	helpers.HttpResponse(w, http.StatusOK)
}

func validateImageToProductDelete(bodyBytes []byte) ([]int64, error) {
	listInt := []int64{}
	userCreateRpcIn := ImageToProductDeleteRpcIn{}

	err := json.Unmarshal(bodyBytes, &userCreateRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return listInt, nil
}
