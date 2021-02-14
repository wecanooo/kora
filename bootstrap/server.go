package bootstrap

import (
	"fmt"
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
	fmt.Printf("\napp runmode is %s, %s\n\n", core.GetConfig().GetMode(), core.GetConfig().String("APP.URL"))
	core.GetApplication().RunServer()
}
