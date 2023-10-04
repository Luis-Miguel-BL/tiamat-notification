package router

import "github.com/Luis-Miguel-BL/tiamat-notification/internal/api"

type MiddlewareFunc interface {
	GetMiddleware() any
}

type Router interface {
	GET(path string, handler api.ControllerFunc, m ...MiddlewareFunc)
	POST(path string, handler api.ControllerFunc, m ...MiddlewareFunc)
	PUT(path string, handler api.ControllerFunc, m ...MiddlewareFunc)
	PATCH(path string, handler api.ControllerFunc, m ...MiddlewareFunc)
	DELETE(path string, handler api.ControllerFunc, m ...MiddlewareFunc)
	Group(prefix string, groupingFn func(group RouteGroup), m ...MiddlewareFunc) RouteGroup
	Use(m ...MiddlewareFunc)
}

type RouteGroup interface {
	Router
}
