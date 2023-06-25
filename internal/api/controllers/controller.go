package controllers

import (
	"context"
	"encoding/json"
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

func ParseRequestBody[T RequestBody](req Request) (parsedBody T, err error) {
	err = json.Unmarshal([]byte(req.Body), parsedBody)
	if err != nil {
		return parsedBody, err
	}

	err = parsedBody.Validate()
	if err != nil {
		return parsedBody, err
	}

	return parsedBody, nil
}

type RequestBody interface {
	Validate() error
}
