package attribute

import (
	"github.com/seivanov1986/gocart/internal/service/attribute"
)

type handle struct {
	service attribute.Service
}

func New(service attribute.Service) *handle {
	return &handle{service: service}
}
