package gocart

import (
	"github.com/seivanov1986/sql_client"

	"github.com/seivanov1986/gocart/external/cache_builder"
	"github.com/seivanov1986/gocart/external/cache_service"
	"github.com/seivanov1986/gocart/external/widget_manager"
	"github.com/seivanov1986/gocart/internal/http/attribute"
	"github.com/seivanov1986/gocart/internal/http/auth"
	"github.com/seivanov1986/gocart/internal/http/category"
	commonHandle "github.com/seivanov1986/gocart/internal/http/common"
	"github.com/seivanov1986/gocart/internal/http/file"
	"github.com/seivanov1986/gocart/internal/http/image_to_category"
	"github.com/seivanov1986/gocart/internal/http/image_to_product"
	"github.com/seivanov1986/gocart/internal/http/page"
	"github.com/seivanov1986/gocart/internal/http/product"
	"github.com/seivanov1986/gocart/internal/http/product_to_category"
	"github.com/seivanov1986/gocart/internal/http/sefurl"
	"github.com/seivanov1986/gocart/internal/http/user"
	auth2 "github.com/seivanov1986/gocart/internal/middleware/auth"
	"github.com/seivanov1986/gocart/internal/middleware/common"
	"github.com/seivanov1986/gocart/internal/middleware/cors"
	"github.com/seivanov1986/gocart/internal/repository"
	attributeService "github.com/seivanov1986/gocart/internal/service/attribute"
	attributeToProductService "github.com/seivanov1986/gocart/internal/service/attribute_to_product"
	authService "github.com/seivanov1986/gocart/internal/service/auth"
	categoryService "github.com/seivanov1986/gocart/internal/service/category"
	commonService "github.com/seivanov1986/gocart/internal/service/common"
	imageToCategoryService "github.com/seivanov1986/gocart/internal/service/image_to_category"
	imageToProductService "github.com/seivanov1986/gocart/internal/service/image_to_product"
	pageService "github.com/seivanov1986/gocart/internal/service/page"
	productService "github.com/seivanov1986/gocart/internal/service/product"
	productToCategoryService "github.com/seivanov1986/gocart/internal/service/product_to_category"
	sefUrlService "github.com/seivanov1986/gocart/internal/service/sefurl"
	user2 "github.com/seivanov1986/gocart/internal/service/user"

	"github.com/seivanov1986/gocart/client"
	"github.com/seivanov1986/gocart/internal/http/attribute_to_product"
	"github.com/seivanov1986/gocart/internal/widget/example"
)

type Options struct {
	database           sql_client.DataBase
	transactionManager sql_client.TransactionManager
	sessionManager     client.SessionManager
	cacheBuilder       client.CacheBuilder
	widgetManager      client.WidgetManager
	buildInWidgets     []string
}

type OptionFunc func(*Options)

func WithDatabase(database sql_client.DataBase) OptionFunc {
	return func(o *Options) {
		o.database = database
	}
}

func WithTransactionManager(trx sql_client.TransactionManager) OptionFunc {
	return func(o *Options) {
		o.transactionManager = trx
	}
}

func WithSessionManager(sessionManager client.SessionManager) OptionFunc {
	return func(o *Options) {
		o.sessionManager = sessionManager
	}
}

func WithCacheBuilder(cacheBuilder client.CacheBuilder) OptionFunc {
	return func(o *Options) {
		o.cacheBuilder = cacheBuilder
	}
}

func WithBuildInWidgets(buildInWidgets []string) OptionFunc {
	return func(o *Options) {
		o.buildInWidgets = buildInWidgets
	}
}

func WithWidgetManager(widgetManager client.WidgetManager) OptionFunc {
	return func(o *Options) {
		o.widgetManager = widgetManager
	}
}

type goCart struct {
	database           sql_client.DataBase
	transactionManager sql_client.TransactionManager
	sessionManager     client.SessionManager
	cacheBuilder       client.CacheBuilder
	widgetManager      client.WidgetManager
	buildInWidgets     []string
}

func New(opts ...OptionFunc) *goCart {
	options := Options{}
	for _, opt := range opts {
		opt(&options)
	}

	return &goCart{
		database:       options.database,
		sessionManager: options.sessionManager,
		cacheBuilder:   options.cacheBuilder,
		widgetManager:  options.widgetManager,
		buildInWidgets: options.buildInWidgets,
	}
}

func (g *goCart) UserHttpHandler() user.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := user2.New(hub)
	return user.New(service)
}

func (g *goCart) AuthHandler() auth.Handle {
	g.checkDatabase()
	g.checkSessionManager()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := authService.New(hub, g.sessionManager)
	return auth.New(service)
}

func (g *goCart) FileHandler() file.Handle {
	return file.New()
}

func (g *goCart) AttributeHandler() attribute.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := attributeService.New(hub)
	return attribute.New(service)
}

func (g *goCart) AttributeToProductHandler() attribute_to_product.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := attributeToProductService.New(hub)
	return attribute_to_product.New(service)
}

func (g *goCart) CategoryHandler() category.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := categoryService.New(hub, g.transactionManager)
	return category.New(service)
}

func (g *goCart) ImageToCategoryHandler() image_to_category.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := imageToCategoryService.New(hub)
	return image_to_category.New(service)
}

func (g *goCart) ImageToProductHandler() image_to_product.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := imageToProductService.New(hub)
	return image_to_product.New(service)
}

func (g *goCart) PageHandler() page.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := pageService.New(hub, g.transactionManager)
	return page.New(service)
}

func (g *goCart) ProductHandler() product.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := productService.New(hub, g.transactionManager)
	return product.New(service)
}

func (g *goCart) ProductToCategoryHandler() product_to_category.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := productToCategoryService.New(hub)
	return product_to_category.New(service)
}

func (g *goCart) SefUrlHandler() sefurl.Handle {
	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	service := sefUrlService.New(hub)
	return sefurl.New(service)
}

func (g *goCart) CommonHandler() commonHandle.Handle {
	service := commonService.New()
	return commonHandle.New(service)
}

func (g *goCart) AuthMiddleware() auth2.Middleware {
	g.checkSessionManager()
	return auth2.New(g.sessionManager)
}

func (g *goCart) CommonMiddleware(serviceBasePath string) common.Middleware {
	return common.New(serviceBasePath)
}

func (g *goCart) CorsMiddleware() cors.Middleware {
	return cors.New()
}

func (g *goCart) checkDatabase() {
	if g.database == nil {
		panic("database must be an object")
	}
}

func (g *goCart) checkTransactionManager() {
	if g.database == nil {
		panic("transaction manager must be an object")
	}
}

func (g *goCart) checkSessionManager() {
	if g.sessionManager == nil {
		panic("session manager must be an object")
	}
}

func (g *goCart) CacheService() cache_service.CacheService {
	if g.cacheBuilder != nil {
		g.cacheBuilder.RegisterWidget("example", example.New())
		return cache_service.New(g.cacheBuilder)
	}

	g.checkDatabase()
	g.checkTransactionManager()

	hub := repository.New(g.database, g.transactionManager)
	if g.widgetManager == nil {
		g.widgetManager = widget_manager.New()
	}

	cacheBuilder := cache_builder.NewBuilder(hub, g.widgetManager)
	return cache_service.New(cacheBuilder)
}
