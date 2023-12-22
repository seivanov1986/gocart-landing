package image_to_category

import (
	"github.com/seivanov1986/gocart/internal/service/image_to_category"
)

type handle struct {
	service image_to_category.Service
}

func New(service image_to_category.Service) *handle {
	return &handle{service: service}
}
