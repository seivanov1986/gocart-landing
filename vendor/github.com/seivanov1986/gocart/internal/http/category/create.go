package category

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/service/category"
)

type CategoryCreateRpcIn struct {
	Name         string  `json:"name"`
	Content      *string `json:"content"`
	Type         int64   `json:"type"`
	Sort         int64   `json:"sort"`
	ShortContent *string `json:"short_content"`
	ImageID      *int64  `json:"id_image"`
	SefURL       string  `json:"sefurl"`
	Template     *string `json:"template"`
	Title        *string `json:"title"`
	Keywords     *string `json:"keywords"`
	Description  *string `json:"description"`
	ParentID     *int64  `json:"id_parent"`
}

type CategoryCreateRpcOut struct {
	ID *int64
}

type CategoryCreateError struct {
	Error string
}

func (u *handle) Create(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 2048))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	CreateListInput, err := validateCategoryCreate(bodyBytes)
	if err != nil {
		fmt.Println(err)
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	userid, err := u.service.Create(r.Context(), *CreateListInput)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError, CategoryCreateError{
			Error: err.Error(),
		})
		return
	}

	helpers.HttpResponse(w, http.StatusOK, CategoryCreateRpcOut{
		ID: userid,
	})
}

func validateCategoryCreate(bodyBytes []byte) (*category.CategoryCreateIn, error) {
	listInt := category.CategoryCreateIn{}
	userCreateRpcIn := CategoryCreateRpcIn{}

	err := json.Unmarshal(bodyBytes, &userCreateRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.Name = userCreateRpcIn.Name
	listInt.Content = userCreateRpcIn.Content
	listInt.Type = userCreateRpcIn.Type
	listInt.Sort = userCreateRpcIn.Sort
	listInt.ShortContent = userCreateRpcIn.ShortContent
	listInt.ImageID = userCreateRpcIn.ImageID
	listInt.Title = userCreateRpcIn.Title
	listInt.Keywords = userCreateRpcIn.Keywords
	listInt.Description = userCreateRpcIn.Description
	listInt.Template = userCreateRpcIn.Template
	listInt.SefURL = userCreateRpcIn.SefURL
	listInt.ParentID = userCreateRpcIn.ParentID

	return &listInt, nil
}
