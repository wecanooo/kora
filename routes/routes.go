package routes

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/wecanooo/kora/core"
	"net/http"
)

// Register 라우팅 핸들러 등록
func Register(router *core.Application) {
	if !core.GetConfig().IsDev() {
		router.Use(middleware.Recover())
	}

	if core.GetConfig().IsDev() {
		router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${status}   ${method}   ${latency_human}               ${uri}\n",
		}))
	}

	router.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))

	router.Pre(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	registerAPI(router)
}