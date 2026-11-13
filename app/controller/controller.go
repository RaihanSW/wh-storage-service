package controller

import (
	"wh-storage-service/app/controller/handler"
	"wh-storage-service/app/controller/handler/middleware"
	"wh-storage-service/app/model/constant"
	"wh-storage-service/config"
	"wh-storage-service/docs"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Controller struct {
	storageHandler       handler.StorageHandlerInterface
	storageStatusHandler handler.StorageStatusHandlerInterface
	warehouseHandler     handler.WarehouseHandlerInterface
	blockLayoutHandler   handler.BlockLayoutHandlerInterface
	generalHandler       handler.GeneralHandlerInterface
	cfg                  *config.ApplicationConfig
	env                  string
}

func NewController(
	storageHandler handler.StorageHandlerInterface,
	storageStatusHandler handler.StorageStatusHandlerInterface,
	warehouseHandler handler.WarehouseHandlerInterface,
	blockLayoutHandler handler.BlockLayoutHandlerInterface,
	generalHandler handler.GeneralHandlerInterface,
	cfg *config.ApplicationConfig,
	env string) *Controller {

	return &Controller{
		storageHandler:       storageHandler,
		storageStatusHandler: storageStatusHandler,
		warehouseHandler:     warehouseHandler,
		blockLayoutHandler:   blockLayoutHandler,
		generalHandler:       generalHandler,
		cfg:                  cfg,
		env:                  env,
	}
}

func (c *Controller) StartRoute() *http.Server {
	r := chi.NewRouter()

	r.Use(middleware.Middleware())

	r.Get("/ping", middleware.SetRequestID(c.generalHandler.Ping))
	r.Get("/health/check", middleware.SetRequestID(c.generalHandler.HealthCheck))
	r.Get("/health/check-db", middleware.SetRequestID(c.generalHandler.HealthCheckDB))

	// Init Swagger with swaggo
	docs.SwaggerInfo.Title = "WH Storage Service"
	docs.SwaggerInfo.Description = fmt.Sprintf("WH Storage Service with Golang - Environtment: %s", strings.ToLower(c.env))
	docs.SwaggerInfo.Version = "1.1"
	switch strings.ToLower(c.env) {
	case constant.EnvLocal:
		docs.SwaggerInfo.Schemes = []string{"http"}
		docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", c.cfg.Port.Service)
		docs.SwaggerInfo.BasePath = "/"
	case constant.EnvDev:
		docs.SwaggerInfo.Schemes = []string{"https"}
		docs.SwaggerInfo.Host = "dc21-dev-api.cinema21.co.id"
		docs.SwaggerInfo.BasePath = "/storage/"
	case constant.EnvStaging:
		docs.SwaggerInfo.Schemes = []string{"https"}
		docs.SwaggerInfo.Host = "dc21-staging-api.cinema21.co.id"
		docs.SwaggerInfo.BasePath = "/storage/"
	}

	if strings.ToLower(c.env) != constant.EnvProd {
		r.Mount("/swagger", httpSwagger.WrapHandler)
	}

	r.Route("/storage", func(r chi.Router) {
		r.Use(middleware.GetLogMiddleware(*c.cfg))
		r.Post("/", c.storageHandler.StorageCreate)
		r.Get("/", c.storageHandler.StorageGetAll)
		r.Get("/{uid}", c.storageHandler.StorageGetByID)
		r.Put("/{uid}", c.storageHandler.StorageUpdate)
		r.Delete("/{uid}", c.storageHandler.StorageDelete)
		r.Route("/status", func(r chi.Router) {
			r.Post("/", c.storageStatusHandler.StorageStatusCreate)
			r.Get("/", c.storageStatusHandler.StorageStatusGetAll)
			r.Get("/{id}", c.storageStatusHandler.StorageStatusGetByID)
			r.Put("/{id}", c.storageStatusHandler.StorageStatusUpdate)
			r.Delete("/{id}", c.storageStatusHandler.StorageStatusDelete)
		})
	})

	r.Route("/warehouse", func(r chi.Router) {
		r.Use(middleware.GetLogMiddleware(*c.cfg))
		r.Post("/", c.warehouseHandler.WarehouseCreate)
		r.Get("/", c.warehouseHandler.WarehouseGetAll)
		r.Get("/{id}", c.warehouseHandler.WarehouseGetByID)
		r.Put("/{id}", c.warehouseHandler.WarehouseUpdate)
		r.Delete("/{id}", c.warehouseHandler.WarehouseDelete)
	})

	r.Route("/block-layout", func(r chi.Router) {
		r.Use(middleware.GetLogMiddleware(*c.cfg))
		r.Post("/", c.blockLayoutHandler.BlockLayoutCreate)
		r.Get("/", c.blockLayoutHandler.BlockLayoutGetAll)
		r.Get("/{id}", c.blockLayoutHandler.BlockLayoutGetByID)
		r.Put("/{id}", c.blockLayoutHandler.BlockLayoutUpdate)
		r.Delete("/{id}", c.blockLayoutHandler.BlockLayoutDelete)
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", c.cfg.Port.Service),
		WriteTimeout: c.cfg.Port.ServiceTimeout * time.Second,
		ReadTimeout:  c.cfg.Port.ServiceTimeout * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return srv
}
