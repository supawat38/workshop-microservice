package routes

import (
	"github.com/gofiber/fiber/v2"

	service_product "app/app/controllers/service_product"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	route := a.Group("/api")

	route.Get("/product", service_product.GetProduct)
	route.Get("/product/:product_code", service_product.GetProductByID)
}
