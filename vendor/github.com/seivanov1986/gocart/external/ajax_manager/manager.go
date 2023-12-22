package ajax_manager

import (
	"net/http"

	"github.com/seivanov1986/gocart/client"
)

type ajaxManager struct {
	handlers   map[string]client.AjaxHandler
}

func New() *ajaxManager {
	return &ajaxManager{
		handlers: map[string]client.AjaxHandler{},
	}
}

func (a *ajaxManager) RegisterPath(name string, widget client.AjaxHandler) {
	a.handlers[name] = widget
}

func (a *ajaxManager) Handler(w http.ResponseWriter, r *http.Request) {
	path := r.Header.Get("x-ajax-path")
	handler, exist := a.handlers[path]
	if !exist {
		return
	}

	result, err := handler.Execute(r.Header, r.Body)
	if err != nil {
		return
	}

	w.Write([]byte(*result))
}
