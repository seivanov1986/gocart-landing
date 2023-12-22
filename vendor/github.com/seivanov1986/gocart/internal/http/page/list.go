package page

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
	"github.com/seivanov1986/gocart/internal/repository/page"
	page2 "github.com/seivanov1986/gocart/internal/service/page"
)

type PageListRpcIn struct {
	Page int64 `json:"page"`
}

type PageListRpcOut struct {
	List  []page.PageListRow
	Total int64
}

func (u *handle) List(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	PageListInput, err := validatePageList(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, http.StatusBadRequest)
		return
	}

	listOut, err := u.service.List(r.Context(), *PageListInput)
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	helpers.HttpResponse(w, http.StatusOK, PageListRpcOut{
		List:  listOut.List,
		Total: listOut.Total,
	})
}

func validatePageList(bodyBytes []byte) (*page2.PageListIn, error) {
	listInt := page2.PageListIn{}
	imageListRpcIn := PageListRpcIn{}

	err := json.Unmarshal(bodyBytes, &imageListRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.Page = imageListRpcIn.Page

	return &listInt, nil
}
