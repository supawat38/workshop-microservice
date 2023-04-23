package routes

import (
	controllersCreatejob "app/app/controllers/module_job"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v2")
	route.Post("/createjob", controllersCreatejob.CreateJob)

}
