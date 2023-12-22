package helpers

import (
	"strings"
)

func GetFileNameByUrl(url string) string {
	return strings.ReplaceAll(url, "/", "x")
}
