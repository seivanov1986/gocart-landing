package image_to_category

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/service/image_to_category"
)

type ImageToCategoryCreateRpcIn struct {
	CategoryID int64 `json:"id_category"`
	ImageID    int64 `json:"id_image"`
}

type ImageToCategoryCreateRpcOut struct {
	ID *int64
}

type ImageToCategoryCreateError struct {
	Error string
}

func (u *handle) Create(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 2048))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	CreateListInput, err := validatePageCreate(bodyBytes)
	if err != nil {
		fmt.Println(err)
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	err = u.service.Create(r.Context(), *CreateListInput)
	if err != nil {
		fmt.Println(err)
		helpers.HttpResponse(w, http.StatusInternalServerError, ImageToCategoryCreateError{
			Error: err.Error(),
		})
		return
	}

	helpers.HttpResponse(w, http.StatusOK)
}

func validatePageCreate(bodyBytes []byte) (*image_to_category.ImageToCategoryCreateIn, error) {
	listInt := image_to_category.ImageToCategoryCreateIn{}
	userCreateRpcIn := ImageToCategoryCreateRpcIn{}

	err := json.Unmarshal(bodyBytes, &userCreateRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.CategoryID = userCreateRpcIn.CategoryID
	listInt.ImageID = userCreateRpcIn.ImageID

	return &listInt, nil
}
