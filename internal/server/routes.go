package server

import (
	customer_controller "github.com/Luis-Miguel-BL/tiamat-notification/internal/api/controller/customer"
	"github.com/Luis-Miguel-BL/tiamat-notification/internal/server/router"
)

func (s *Server) setupApiRoutes(
	r router.Router,
	saveCustomerController *customer_controller.SaveCustomerController,
	createCustomerEventController *customer_controller.CreateCustomerEventController,
) {
	r.POST("/identify", saveCustomerController.Execute)
	r.POST("/tracking", createCustomerEventController.Execute)
}
