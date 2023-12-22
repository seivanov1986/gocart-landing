package cache_builder

import (
	"context"

	"github.com/seivanov1986/gocart/client"
)

type Service interface {
	Make(ctx context.Context) error
}

type WidgetManager interface {
	Render(ctx context.Context, name string) (*string, error)
	Register(name string, widget client.Widget)
	SetAssets(assetManager client.AssetManager)
}

type Widget interface {
	Execute() (*string, error)
}

type CacheBuilder interface {
	Pages(ctx context.Context) ([]client.UrlListRow, error)
	Handler(ctx context.Context, pages []client.UrlListRow) error
	RegisterWidget(name string, widget Widget)
}
