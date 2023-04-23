package routes

import (
	"github.com/gofiber/fiber/v2"

	service_authentication "app/app/controllers/service_authentication"
	service_order "app/app/controllers/service_order"
	service_product "app/app/controllers/service_product"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api")

	//เข้าสู่ระบบ
	route.Post("/login", service_authentication.Login)
	route.Post("/logout", service_authentication.Logout)

	//สมัครสมาชิก
	route.Post("/register", service_authentication.Register)

	//เพิ่มสินค้า
	route.Post("/create_product", service_product.CreateProduct)

	//สร้างใบสั้งซื้อ & ยกเลิกใบสั้งซื้อ
	route.Post("/create_purchaseorder", service_order.CreatePurchaseorder)
	route.Post("/cancel_purchaseorder", service_order.CancelPurchaseorder)

}
