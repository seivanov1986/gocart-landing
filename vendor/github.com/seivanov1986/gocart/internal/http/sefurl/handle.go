package sefurl

import (
	"github.com/seivanov1986/gocart/internal/service/sefurl"
)

type handle struct {
	sefUrlService sefurl.Service
}

func New(sefUrlService sefurl.Service) *handle {
	return &handle{sefUrlService: sefUrlService}
}
