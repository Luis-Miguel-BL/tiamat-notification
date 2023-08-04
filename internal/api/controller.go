package api

import (
	"context"
)

type Controller interface {
	Execute(ctx context.Context, request Request) (response Response)
}

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
