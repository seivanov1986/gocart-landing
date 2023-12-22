package image_to_category

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	image_to_category2 "github.com/seivanov1986/gocart/internal/repository/image_to_category"
	"github.com/seivanov1986/gocart/internal/service/image_to_category"
)

type ImageToCategoryListRpcIn struct {
	Page       int64 `json:"page"`
	CategoryID int64 `json:"id_category"`
}

type ImageToCategoryListRpcOut struct {
	List  []image_to_category2.ImageToCategoryListRow
	Total int64
}

func (u *handle) List(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	PageListInput, err := validatePageListV2(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	listOut, err := u.service.List(r.Context(), *PageListInput)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK, ImageToCategoryListRpcOut{
		List:  listOut.List,
		Total: listOut.Total,
	})
}

func validatePageListV2(bodyBytes []byte) (*image_to_category.ImageToCategoryListIn, error) {
	listInt := image_to_category.ImageToCategoryListIn{}
	imageListRpcIn := ImageToCategoryListRpcIn{}

	err := json.Unmarshal(bodyBytes, &imageListRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.CategoryID = imageListRpcIn.CategoryID
	listInt.Page = imageListRpcIn.Page

	return &listInt, nil
}
