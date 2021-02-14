package context

import (
	"github.com/labstack/echo/v4"
	"strings"
)

type AppContext struct {
	echo.Context
}

type (
	AppHandlerFunc = func(c *AppContext) error
	AppRegisterFunc = func(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
)

// NewAppContext application context instance 생성
func NewAppContext(context echo.Context) *AppContext {
	return &AppContext{Context: context}
}

// RegisterHandler route handler 등록
func RegisterHandler(fn AppRegisterFunc, path string, h AppHandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	if path != "" && !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return fn(path, func(c echo.Context) error {
		cc, ok := c.(*AppContext)
		if !ok {
			cc = NewAppContext(c)
			return h(cc)
		}
		return h(cc)
	}, m...)
}