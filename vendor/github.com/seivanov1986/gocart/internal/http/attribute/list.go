package attribute

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	attribute2 "github.com/seivanov1986/gocart/internal/service/attribute"

	"github.com/seivanov1986/gocart/internal/repository/attribute"
)

type AttributeListRpcIn struct {
	Page int64 `json:"page"`
}

type AttributeListRpcOut struct {
	List  []attribute.AttributeListRow
	Total int64
}

func (u *handle) List(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	AttributeListInput, err := validateAttributeList(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	listOut, err := u.service.List(r.Context(), *AttributeListInput)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK, AttributeListRpcOut{
		List:  listOut.List,
		Total: listOut.Total,
	})
}

func validateAttributeList(bodyBytes []byte) (*attribute2.AttributeListIn, error) {
	listInt := attribute2.AttributeListIn{}
	attributeListRpcIn := AttributeListRpcIn{}

	err := json.Unmarshal(bodyBytes, &attributeListRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.Page = attributeListRpcIn.Page

	return &listInt, nil
}
