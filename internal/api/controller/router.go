package controller

import "net/http"

type middlewareFunc func(context.Context, Request) (context.Context, error)

type Route struct {
	r map[string]map[string] ControllerFunc
	middleware []middlewareFunc
}
type RouterGroup struct {
	rg map[string]Route
	middleware []middlewareFunc
}
type Router struct {
	r map[string]
}