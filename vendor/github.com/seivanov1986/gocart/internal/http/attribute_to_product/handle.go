package attribute_to_product

import (
	"github.com/seivanov1986/gocart/internal/service/attribute_to_product"
)

type handle struct {
	service attribute_to_product.Service
}

func New(service attribute_to_product.Service) *handle {
	return &handle{service: service}
}
