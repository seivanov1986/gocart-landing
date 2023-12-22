package helpers

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func SaveFile(path string, bodyReader io.Reader) (err error) {
	replacer := strings.NewReplacer("../", "", "..\\", "")
	path = replacer.Replace(path)
	baseDir := filepath.Dir(path)

	err = os.MkdirAll(baseDir, 0777)
	if err != nil {
		return err
	}

	w, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		err = w.Close()
	}()

	_, err = io.Copy(w, bodyReader)
	if err != nil {
		return err
	}

	return nil
}
