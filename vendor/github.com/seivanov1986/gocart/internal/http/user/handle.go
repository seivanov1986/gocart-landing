package user

import (
	userService "github.com/seivanov1986/gocart/internal/service/user"
)

type handle struct {
	service userService.Service
}

func New(service userService.Service) *handle {
	return &handle{
		service: service,
	}
}
