package product

import (
	"github.com/seivanov1986/gocart/internal/service/product"
)

type handle struct {
	service product.Service
}

func New(service product.Service) *handle {
	return &handle{service: service}
}
