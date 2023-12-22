package attribute

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/service/attribute"
)

type AttributeCreateRpcIn struct {
	Name      string  `json:"name"`
	Signature *string `json:"signature"`
}

type AttributeCreateRpcOut struct {
	ID *int64
}

type AttributeCreateError struct {
	Error string
}

func (u *handle) Create(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	CreateListInput, err := validateAttributeCreate(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	userid, err := u.service.Create(r.Context(), *CreateListInput)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError, AttributeCreateError{
			Error: err.Error(),
		})
		return
	}

	helpers.HttpResponse(w, http.StatusOK, AttributeCreateRpcOut{
		ID: userid,
	})
}

func validateAttributeCreate(bodyBytes []byte) (*attribute.AttributeCreateIn, error) {
	listInt := attribute.AttributeCreateIn{}
	userCreateRpcIn := AttributeCreateRpcIn{}

	err := json.Unmarshal(bodyBytes, &userCreateRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.Name = userCreateRpcIn.Name
	listInt.Signature = userCreateRpcIn.Signature

	return &listInt, nil
}
