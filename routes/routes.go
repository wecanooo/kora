package routes

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/wecanooo/kora/core"
	"net/http"

	mw "github.com/wecanooo/kora/routes/middleware"
)

// Register register middlewares
func Register(router *core.Application) {
	if !core.GetConfig().IsDev() {
		router.Use(middleware.Recover())
	}

	if core.GetConfig().IsDev() {
		router.Use(middleware.Logger())
	} else {
		router.Use(mw.ZapLogger(core.GetLog()))
	}

	router.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))

	router.Pre(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	registerAPI(router)
}