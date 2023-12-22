package cache_builder

import (
	"bytes"
	"context"
	"regexp"
	"strings"
	"text/template"

	"github.com/seivanov1986/gocart/client"
	"github.com/seivanov1986/gocart/external/asset_manager"
	"github.com/seivanov1986/gocart/external/observer"
	"github.com/seivanov1986/gocart/helpers"
)

type builder struct {
	widgetManager WidgetManager
}

func (b *builder) Pages(ctx context.Context) ([]client.UrlListRow, error) {
	return []client.UrlListRow{}, nil
}

func NewBuilder(widgetManager WidgetManager) *builder {
	return &builder{widgetManager: widgetManager}
}

func (b *builder) RegisterWidget(name string, widget client.Widget) {
	b.widgetManager.Register(name, widget)
}

func (b *builder) Handler(ctx context.Context, pages []client.UrlListRow) error {
	serviceBasePath := observer.GetServiceBasePath(ctx)

	templateFiles := []string{
		serviceBasePath + "/schemes/templates/layouts/common.html",
	}

	tmpl, err := template.New("common").ParseFiles(templateFiles...)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}

	assetManager := asset_manager.New()
	b.widgetManager.SetAssets(assetManager)

	err = tmpl.ExecuteTemplate(buf, "common", nil)
	if err != nil {
		return err
	}

	content := buf.String()

	reg, _ := regexp.Compile(`{#outertemplate%([A-Za-z_0-9]+)#}`)
	for _, match := range reg.FindAllStringSubmatch(content, -1) {
		match_1 := match[1]
		if len(match) > 2 {
			match_1 = match[2]
		}

		res, _ := b.widgetManager.Render(ctx, match_1)
		if res != nil {
			content = strings.Replace(content, match[0], *res, -1)
		}
	}

	fileName := helpers.GetFileNameByUrl("/")
	return helpers.SaveFile("/tmp/cache/"+fileName, strings.NewReader(content))
}
