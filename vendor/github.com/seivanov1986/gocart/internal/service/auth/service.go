package user

import (
	"github.com/seivanov1986/gocart/internal/repository"
)

type service struct {
	hub            repository.Hub
	sessionManager SessionManager
}

func New(hub repository.Hub, sessionManager SessionManager) *service {
	return &service{hub: hub, sessionManager: sessionManager}
}
