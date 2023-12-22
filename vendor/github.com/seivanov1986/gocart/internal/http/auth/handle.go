package auth

import (
	auth "github.com/seivanov1986/gocart/internal/service/auth"
)

type handle struct {
	service auth.Service
}

func New(service auth.Service) *handle {
	return &handle{
		service: service,
	}
}
