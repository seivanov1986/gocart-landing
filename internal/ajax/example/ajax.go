package example

import (
	"io"
	"net/http"
)

type loggerPlugin struct {
}

func New() *loggerPlugin {
	return &loggerPlugin{}
}

func (l *loggerPlugin) Execute(header http.Header, body io.ReadCloser) (*string, error) {
	result := "example"
	return &result, nil
}
