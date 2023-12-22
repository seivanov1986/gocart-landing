package attribute

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	attribute2 "github.com/seivanov1986/gocart/internal/service/attribute"
)

type AttributeUpdateRpcIn struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Signature *string `json:"signature"`
}

type AttributeUpdateError struct {
	Error string
}

func (u *handle) Update(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	CreateListInput, err := validateAttributeUpdate(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	err = u.service.Update(r.Context(), *CreateListInput)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError, AttributeCreateError{
			Error: err.Error(),
		})
		return
	}

	helpers.HttpResponse(w, http.StatusOK)
}

func validateAttributeUpdate(bodyBytes []byte) (*attribute2.AttributeUpdateInput, error) {
	listInt := attribute2.AttributeUpdateInput{}
	userCreateRpcIn := AttributeUpdateRpcIn{}

	err := json.Unmarshal(bodyBytes, &userCreateRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.ID = userCreateRpcIn.ID
	listInt.Name = userCreateRpcIn.Name
	listInt.Signature = userCreateRpcIn.Signature

	return &listInt, nil
}
