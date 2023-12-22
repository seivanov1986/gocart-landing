package cache_service

import (
	"context"
	"github.com/seivanov1986/gocart/client"
)

type CacheService interface {
	Make(ctx context.Context)
}

type cacheService struct {
	cacheBuilder client.CacheBuilder
}

func New(cacheBuilder client.CacheBuilder) *cacheService {
	return &cacheService{cacheBuilder: cacheBuilder}
}

func (c *cacheService) Make(ctx context.Context) {
	pages, _ := c.cacheBuilder.Pages(ctx)
	c.cacheBuilder.Handler(ctx, pages)
}
