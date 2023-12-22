package page

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	page2 "github.com/seivanov1986/gocart/internal/service/page"
)

type PageUpdateRpcIn struct {
	ID           int64   `json:"id"`
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
}

type PageUpdateRpcOut struct {
	ID *int64
}

type PageUpdateError struct {
	Error string
}

func (u *handle) Update(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 2048))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	CreateListInput, err := validatePageUpdate(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	err = u.service.Update(r.Context(), *CreateListInput)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError, PageCreateError{
			Error: err.Error(),
		})
		return
	}

	helpers.HttpResponse(w, http.StatusOK)
}

func validatePageUpdate(bodyBytes []byte) (*page2.PageUpdateInput, error) {
	listInt := page2.PageUpdateInput{}
	userCreateRpcIn := PageUpdateRpcIn{}

	err := json.Unmarshal(bodyBytes, &userCreateRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.ID = userCreateRpcIn.ID
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

	return &listInt, nil
}
