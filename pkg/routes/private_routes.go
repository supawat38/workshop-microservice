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

	//เข้าสู่ระบบ (ไม่ต้องใช้ Token) && สมัครสมาชิก (ไม่ต้องใช้ Token)
	services_authen := a.Group("/api")
	services_authen.Post("/login", service_authentication.Login)
	services_authen.Post("/logout", middleware.JWTProtected(), service_authentication.Logout)
	services_authen.Post("/register", service_authentication.Register)

	//เพิ่มสินค้า
	services_product := a.Group("/api")
	services_product.Post("/create_product", middleware.JWTProtected(), service_product.CreateProduct)

	//สร้างใบสั้งซื้อ && ยกเลิกใบสั้งซื้อ
	services_order := a.Group("/api")
	services_order.Post("/create_purchaseorder", middleware.JWTProtected(), service_order.CreatePurchaseorder)
	services_order.Post("/cancel_purchaseorder", middleware.JWTProtected(), service_order.CancelPurchaseorder)

}
