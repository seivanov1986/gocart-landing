package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/seivanov1986/gocart"
	"github.com/seivanov1986/gocart/external/ajax_manager"
	"github.com/seivanov1986/gocart/external/observer"
	"github.com/seivanov1986/gocart/external/widget_manager"

	ajaxExample "github.com/seivanov1986/gocart-landing/internal/ajax/example"
	"github.com/seivanov1986/gocart-landing/internal/widget/example"
	"github.com/seivanov1986/gocart-landing/pkg/cache_builder"
)

func main() {
	serviceBasePath := os.Getenv("SERVICE_BASE_PATH")
	if serviceBasePath == "" {
		panic("service base path is not found")
	}

	router := mux.NewRouter()
	ctx := context.Background()
	ctx = observer.SetServiceBasePath(ctx, serviceBasePath)

	widgetManger := widget_manager.New()
	ajaxManager := ajax_manager.New()
	ajaxManager.RegisterPath("example", ajaxExample.New())
	widgetManger.Register("exampleout", example.New())
	cacheBuilder := cache_builder.NewBuilder(widgetManger)

	goLib := gocart.New(
		gocart.WithCacheBuilder(cacheBuilder),
	)

	cacheService := goLib.CacheService()
	cacheService.Make(ctx)

	corsMiddleware := goLib.CorsMiddleware()
	commonMiddleware := goLib.CommonMiddleware(serviceBasePath)

	commonHandle := goLib.CommonHandler()

	router.Use(commonMiddleware.Handle, corsMiddleware.Handle)
	notFoundHandler := commonMiddleware.Wrapper(commonHandle.Process)
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	router.HandleFunc("/ajax", ajaxManager.Handler).
		Methods(http.MethodPost, http.MethodOptions)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("ready")
	log.Fatal(srv.ListenAndServe())
}
