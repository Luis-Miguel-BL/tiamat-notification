package adapters

import (
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/api"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/server/router"
	"github.com/labstack/echo/v4"
)

type echoRouterAdapter struct {
	echo *echo.Group
}

func NewEchoRouterAdapter(e *echo.Echo) router.Router {
	return &echoRouterAdapter{
		echo: e.Group(""),
	}
}

func (a *echoRouterAdapter) GET(path string, handler api.ControllerFunc, m ...router.MiddlewareFunc) {
	a.echo.GET(a.adaptPath(path), a.adaptEchoRoute(handler), a.adaptEchoMiddlewares(m)...)
}

func (a *echoRouterAdapter) POST(path string, handler api.ControllerFunc, m ...router.MiddlewareFunc) {
	a.echo.POST(a.adaptPath(path), a.adaptEchoRoute(handler), a.adaptEchoMiddlewares(m)...)
}

func (a *echoRouterAdapter) PUT(path string, handler api.ControllerFunc, m ...router.MiddlewareFunc) {
	a.echo.PUT(a.adaptPath(path), a.adaptEchoRoute(handler), a.adaptEchoMiddlewares(m)...)
}

func (a *echoRouterAdapter) PATCH(path string, handler api.ControllerFunc, m ...router.MiddlewareFunc) {
	a.echo.PATCH(a.adaptPath(path), a.adaptEchoRoute(handler), a.adaptEchoMiddlewares(m)...)
}

func (a *echoRouterAdapter) DELETE(path string, handler api.ControllerFunc, m ...router.MiddlewareFunc) {
	a.echo.DELETE(a.adaptPath(path), a.adaptEchoRoute(handler), a.adaptEchoMiddlewares(m)...)
}

func (a *echoRouterAdapter) Use(m ...router.MiddlewareFunc) {
	a.echo.Use(a.adaptEchoMiddlewares(m)...)
}

func (a *echoRouterAdapter) Group(prefix string, groupingFn func(group router.RouteGroup), m ...router.MiddlewareFunc) router.RouteGroup {
	echoGroup := a.echo.Group(prefix, a.adaptEchoMiddlewares(m)...)
	groupRouter := &echoRouterAdapter{
		echo: echoGroup,
	}

	if groupingFn != nil {
		groupingFn(groupRouter)
	}

	return groupRouter
}

func (a *echoRouterAdapter) adaptPath(path string) string {
	if path == "/" {
		return ""
	}
	return path
}

func (a *echoRouterAdapter) adaptEchoRoute(handler api.ControllerFunc) func(c echo.Context) error {
	return func(c echo.Context) error {
		request := a.adaptEchoRequest(c)
		response := handler(c.Request().Context(), request)

		return c.JSON(response.StatusCode, response.Body)
	}
}

func (a *echoRouterAdapter) adaptEchoRequest(context echo.Context) api.Request {
	return NewEchoRequestAdapter(context)
}

func (a *echoRouterAdapter) adaptEchoMiddlewares(handlers []router.MiddlewareFunc) []echo.MiddlewareFunc {
	m := []echo.MiddlewareFunc{}
	for _, handler := range handlers {
		m = append(m, handler.GetMiddleware().(echo.MiddlewareFunc))
	}
	return m
}
