package echo

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/mimani68/fintech-core/config"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
	db "github.com/tendermint/tm-db"
)

// @title Swagger API
// @version 1.0
// @description Conduit API
// @title Conduit API

// @host 0.0.0.0:3000
// @BasePath /api

// @schemes http https
// @produce	application/json
// @consumes application/json

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func NewEchoApp(db *db.DB, config config.Config, router echo.Router) {
	r := router.New()
	routers := r.Group("/api")

	r.GET("/swagger/*", echoSwagger.WrapHandler)

	d := db.New()
	db.AutoMigrate(d)

	fo := service.NewFormService(d)
	h := handler.NewHandler(fo)
	h.RegisterRouters(routers)

	r.Logger.Fatal(r.Start(fmt.Sprintf("%s:%s", config.ADDRESS, config.PORT)))
}
