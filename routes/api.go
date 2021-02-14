package routes

import (
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
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

	router.Group(APIPrefix, middleware.CORS())

	//user := e.Group("/users")
	//{
	//
	//}
}