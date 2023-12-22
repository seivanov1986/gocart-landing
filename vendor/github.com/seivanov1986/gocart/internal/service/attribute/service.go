package attribute

import (
	repository2 "github.com/seivanov1986/gocart/internal/repository"
)

type service struct {
	hub repository2.Hub
}

func New(hub repository2.Hub) *service {
	return &service{hub: hub}
}
