package adapters

import (
	"io"

	"github.com/Luis-Miguel-BL/tiamat-notification/internal/api"
	"github.com/labstack/echo/v4"
)

func NewEchoRequestAdapter(echo echo.Context) api.Request {
	body, _ := io.ReadAll(echo.Request().Body)
	return api.Request{
		Method: echo.Request().Method,
		Body:   string(body),
	}
}
