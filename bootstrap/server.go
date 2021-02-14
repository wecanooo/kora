package bootstrap

import (
	"github.com/labstack/echo/v4"
	"github.com/wecanooo/kora/core"
)

func SetupServer() {
	core.SetupLog()

	e := echo.New()
	e.Debug = core.GetConfig().IsDev()
	e.HideBanner = true

	core.NewApplication(e)
}

func RunServer() {
	core.GetApplication().RunServer()
}
