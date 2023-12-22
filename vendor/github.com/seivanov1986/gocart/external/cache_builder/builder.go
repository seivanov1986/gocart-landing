package cache_builder

import (
	"context"

	"github.com/seivanov1986/gocart/client"
	"github.com/seivanov1986/gocart/internal/repository"
)

type builder struct {
	hub repository.Hub
	widgetManager client.WidgetManager
	assetManager client.AssetManager
}

func NewBuilder(hub repository.Hub, widgetManager client.WidgetManager) *builder {
	return &builder{hub: hub, widgetManager: widgetManager}
}

func (b *builder) RegisterWidget(name string, widget client.Widget) {
	b.widgetManager.Register(name, widget)
}

func (b *builder) SetAssets(assetManager client.AssetManager) {
	b.assetManager = assetManager
}

func (b *builder) Pages(ctx context.Context) ([]client.UrlListRow, error) {
	return nil, nil
}

func (b *builder) Handler(ctx context.Context, pages []client.UrlListRow) error {
	return nil
}
