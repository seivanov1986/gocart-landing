package file

import (
	"net/http"
)

func (c *handle) Static(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	serviceBasePath := ctx.Value(serviceBasePathKey).(string)
	httpFileServerHandler := http.FileServer(http.Dir(serviceBasePath + fileServerHandlerPostfix))
	http.StripPrefix(stripPrefix, httpFileServerHandler).ServeHTTP(w, r)
}
