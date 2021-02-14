package bootstrap

import (
	"github.com/labstack/echo/v4"
	"github.com/wecanooo/kora/core"
	"github.com/wecanooo/kora/routes"
)

func SetupServer() {
	core.SetupLog()

	e := echo.New()
	e.Debug = core.GetConfig().IsDev()
	e.HideBanner = true

	core.NewApplication(e)
	core.GetApplication().RegisterRoutes(routes.Register)
	core.GetApplication().PrintRoutes(core.GetConfig().String("APP.TEMP_DIR") + "/routes.json")
}

func RunServer() {
	core.GetApplication().RunServer()
}
