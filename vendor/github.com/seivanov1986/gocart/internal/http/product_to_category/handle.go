package product_to_category

import (
	"github.com/seivanov1986/gocart/internal/service/product_to_category"
)

type handle struct {
	service product_to_category.Service
}

func New(service product_to_category.Service) *handle {
	return &handle{service: service}
}
