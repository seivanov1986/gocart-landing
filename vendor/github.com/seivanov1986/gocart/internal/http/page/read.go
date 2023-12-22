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

type PageReadRpcIn struct {
	ID int64 `json:"id"`
}

type PageReadRpcOut struct {
	Row page.PageReadRow
}

func (u *handle) Read(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, 100))
	if err != nil {
		helpers.HttpResponse(w, http.StatusInternalServerError)
		return
	}

	UserRow, err := validatePageRead(bodyBytes)
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

	helpers.HttpResponse(w, http.StatusOK, PageReadRpcOut{
		Row: *listOut,
	})
}

func validatePageRead(bodyBytes []byte) (*page2.PageReadIn, error) {
	listInt := page2.PageReadIn{}
	userReadRpcIn := PageReadRpcIn{}

	err := json.Unmarshal(bodyBytes, &userReadRpcIn)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	listInt.ID = userReadRpcIn.ID

	return &listInt, nil
}
