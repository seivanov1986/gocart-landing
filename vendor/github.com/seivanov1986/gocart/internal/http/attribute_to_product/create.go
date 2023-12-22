package attribute_to_product

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	attribute_to_product2 "github.com/seivanov1986/gocart/internal/service/attribute_to_product"
)

type AttributeToProductCreateRpcIn struct {
}

type AttributeToProductCreateRpcOut struct {
}

type AttributeToProductCreateError struct {
}

func (u *handle) Create(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	listInput, err := validateAttributeToProductCreate(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	err = u.service.Create(r.Context(), *listInput)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError, AttributeToProductCreateError{})
		return
	}

	helpers.HttpResponse(w, http.StatusOK, AttributeToProductCreateRpcOut{})
}

func validateAttributeToProductCreate(bodyBytes []byte) (*attribute_to_product2.AttributeToProductCreateInput, error) {
	var (
		out             attribute_to_product2.AttributeToProductCreateInput
		userCreateRpcIn AttributeToProductCreateRpcIn
	)

	if err := json.Unmarshal(bodyBytes, &userCreateRpcIn); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &out, nil
}
