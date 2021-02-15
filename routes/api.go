package routes

import (
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/wecanooo/kora/app/controllers/api"
	"github.com/wecanooo/kora/app/services"
	"github.com/wecanooo/kora/core"
)

const (
	// APIPrefix api prefix
	APIPrefix = "/api"
)

func registerAPI(router *core.Application) {
	if core.GetConfig().IsDev() {
		router.GET("/api-doc/*", echoSwagger.WrapHandler).Name = "api-doc"
	}

	e := router.Group(APIPrefix, middleware.CORS())
	v1 := e.Group("/v1")

	user := v1.Group("/users")
	{
		uc := api.NewUserController(services.NewUserServices())
		router.RegisterHandler(user.GET, "", uc.Index).Name = "user.index"
		router.RegisterHandler(user.POST, "", uc.Create).Name = "user.create"
	}
}