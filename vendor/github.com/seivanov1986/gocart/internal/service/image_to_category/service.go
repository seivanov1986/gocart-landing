package image_to_category

import (
	"github.com/seivanov1986/gocart/internal/repository"
)

type service struct {
	hub repository.Hub
}

func New(hub repository.Hub) *service {
	return &service{hub: hub}
}
