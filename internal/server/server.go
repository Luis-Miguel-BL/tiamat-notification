package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	adapters "github.com/Luis-Miguel-BL/tiamat-notification/internal/adapter"
	controller "github.com/Luis-Miguel-BL/tiamat-notification/internal/api/controller/customer"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/application/usecase"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/config"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/infra/logger"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo           *echo.Echo
	config         *config.Config
	log            logger.Logger
	usecaseManager usecase.UsecaseManager
}

func NewServer(log logger.Logger, cfg *config.Config, usecaseManager usecase.UsecaseManager) (*Server, error) {
	server := &Server{
		echo:           echo.New(),
		log:            log,
		config:         cfg,
		usecaseManager: usecaseManager,
	}

	server.setup()
	return server, nil
}

func (s *Server) setup() {
	s.echo.HideBanner = true

	s.echo.GET("/health/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	s.setupApiRoutes(
		adapters.NewEchoRouterAdapter(s.echo),
		controller.NewSaveCustomerController(s.usecaseManager.SaveCustomer, s.log),
		controller.NewCreateCustomerEventController(s.usecaseManager.CreateCustomerEvent, s.log),
	)
}

func (s *Server) waitForShutdown(ctx context.Context) {
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.echo.Shutdown(ctx); err != nil {
		s.echo.Logger.Fatal(err)
	}
}

func (s *Server) Run(ctx context.Context) {
	go func() {
		address := fmt.Sprintf(":%d", s.config.Server.Port)
		if err := s.echo.Start(address); err != nil && err != http.ErrServerClosed {
			s.echo.Logger.Fatal(err)
		}
	}()
	s.waitForShutdown(ctx)
}
