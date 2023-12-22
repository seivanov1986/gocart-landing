package category

import (
	"github.com/seivanov1986/gocart/internal/service/category"
)

type handle struct {
	service category.Service
}

func New(service category.Service) *handle {
	return &handle{service: service}
}
