package product

import (
	"github.com/seivanov1986/sql_client"

	repository2 "github.com/seivanov1986/gocart/internal/repository"
)

const (
	productType = 3
)

type service struct {
	hub       repository2.Hub
	TrManager sql_client.TransactionManager
}

func New(hub repository2.Hub, TrManager sql_client.TransactionManager) *service {
	return &service{hub: hub, TrManager: TrManager}
}
