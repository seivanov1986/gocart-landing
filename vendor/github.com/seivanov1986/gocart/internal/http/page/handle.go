package page

import (
	"github.com/seivanov1986/gocart/internal/service/page"
)

type handle struct {
	service page.Service
}

func New(service page.Service) *handle {
	return &handle{service: service}
}
