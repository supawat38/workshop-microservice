package routes

import (
	"github.com/gofiber/fiber/v2"

	service_authentication "app/app/controllers/service_authentication"
	service_order "app/app/controllers/service_order"
	service_product "app/app/controllers/service_product"
	"app/pkg/middleware"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api")

	//เข้าสู่ระบบ (ไม่ต้องใช้ Token)
	route.Post("/login", service_authentication.Login)

	//ออกจากระบบ
	route.Post("/logout", middleware.JWTProtected(), service_authentication.Logout)

	//สมัครสมาชิก (ไม่ต้องใช้ Token)
	route.Post("/register", service_authentication.Register)

	//เพิ่มสินค้า
	route.Post("/create_product", middleware.JWTProtected(), service_product.CreateProduct)

	//สร้างใบสั้งซื้อ
	route.Post("/create_purchaseorder", middleware.JWTProtected(), service_order.CreatePurchaseorder)

	//ยกเลิกใบสั้งซื้อ
	route.Post("/cancel_purchaseorder", middleware.JWTProtected(), service_order.CancelPurchaseorder)

}
