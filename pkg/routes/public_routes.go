package routes

import (
	controllersCreatejob "app/app/controllers/module_job"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v2")

	route.Get("/listjob", controllersCreatejob.ListJob)
}
