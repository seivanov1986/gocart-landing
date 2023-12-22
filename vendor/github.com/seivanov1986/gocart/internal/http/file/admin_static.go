package file

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/seivanov1986/gocart/helpers"
)

func (c *handle) AdminStatic(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	serviceBasePath := ctx.Value(serviceBasePathKey).(string)

	if !strings.HasPrefix(r.URL.Path, adminPrefix) {
		fileReader, err := os.Open(serviceBasePath + adminPostfix)
		if err != nil {
			helpers.HttpResponse(w, http.StatusInternalServerError)
		}

		bytes, err := io.ReadAll(fileReader)
		if err != nil {
			helpers.HttpResponse(w, http.StatusInternalServerError)
		}

		_, err = w.Write(bytes)
		if err != nil {
			helpers.HttpResponse(w, http.StatusInternalServerError)
		}
	} else {
		httpFileServerHandler := http.FileServer(http.Dir(serviceBasePath + serviceBasePathPostfix))
		http.StripPrefix("/admin", httpFileServerHandler).ServeHTTP(w, r)
	}
}
