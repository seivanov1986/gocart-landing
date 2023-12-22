package file

import (
	"net/http"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func (c *handle) Dynamic(w http.ResponseWriter, r *http.Request) {
	reg, _ := regexp.Compile(`/dynamic/([0-9]+)x([0-9]+)/(.+)`)
	matches := reg.FindAllStringSubmatch(r.URL.Path, 1)
	if len(matches) == 0 {
		return
	}

	if len(matches[0]) != 4 {
		return
	}

	originalPath := path.Dir(matches[0][3])
	originalFileName := path.Base(matches[0][3])
	fileNameExt := filepath.Ext(originalFileName)
	fileName := strings.TrimSuffix(originalFileName, fileNameExt)
	fileName = fileName + "_" + matches[0][1] + "x" + matches[0][2] + fileNameExt
	r.URL.Path = fileName

	httpFileServerHandler := http.FileServer(http.Dir(dynamicPrefix + originalPath))
	http.StripPrefix("", httpFileServerHandler).ServeHTTP(w, r)
}
