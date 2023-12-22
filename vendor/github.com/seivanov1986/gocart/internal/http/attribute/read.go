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

type AttributeReadRpcIn struct {
	ID int64 `json:"id"`
}

type AttributeReadRpcOut struct {
	Row attribute.AttributeReadRow
}

func (u *handle) Read(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	UserRow, err := validateAttributeRead(bodyBytes)
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

	helpers.HttpResponse(w, http.StatusOK, AttributeReadRpcOut{
		Row: *listOut,
	})
}

func validateAttributeRead(bodyBytes []byte) (*attribute2.AttributeReadIn, error) {
	listInt := attribute2.AttributeReadIn{}
	userReadRpcIn := AttributeReadRpcIn{}

	err := json.Unmarshal(bodyBytes, &userReadRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.ID = userReadRpcIn.ID

	return &listInt, nil
}
