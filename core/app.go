package core

import "github.com/labstack/echo/v4"

type Application struct {
	*echo.Echo
}

// NewApplication server instance 생성
func NewApplication(echo *echo.Echo) {
	application = &Application{Echo: echo}
}

// RunServer server 시작
func (app *Application) RunServer() {
	app.Logger.Fatal(app.Start(GetConfig().DefaultString("APP.ADDR", ":3000")))
}