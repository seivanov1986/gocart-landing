package common

import (
	"github.com/seivanov1986/gocart/internal/service/common"
)

type handle struct {
	service common.Service
}

func New(service common.Service) *handle {
	return &handle{
		service: service,
	}
}
