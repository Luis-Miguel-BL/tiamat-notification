package api

import "context"

type Request struct {
	Method string
	Body   string
}
type Response struct {
	StatusCode int
	Body       string
}

type RequestBody interface {
	Validate() error
}

type ControllerFunc func(context.Context, Request) Response
