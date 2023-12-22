package common

import (
	"context"
	"io"
	"os"

	"github.com/seivanov1986/gocart/helpers"
)

func (s *service) Process(ctx context.Context, path string) ([]byte, error) {
	fileName := helpers.GetFileNameByUrl(path)

	fileReader, err := os.Open("/tmp/cache/" + fileName)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(fileReader)
}
