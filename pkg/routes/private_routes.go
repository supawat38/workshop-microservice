package routes

import (
	"github.com/gofiber/fiber/v2"

	service_authentication "app/app/controllers/service_authentication"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api")

	//เข้าสู่ระบบ
	route.Post("/login", service_authentication.Login)
	route.Post("/logout", service_authentication.Logout)

	//สมัครสมาชิก
	route.Post("/register", service_authentication.Register)
}
