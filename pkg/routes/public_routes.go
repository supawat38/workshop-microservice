package routes

import (
	"github.com/gofiber/fiber/v2"

	service_order "app/app/controllers/service_order"
	service_product "app/app/controllers/service_product"
	service_member "app/app/controllers/service_user"
	"app/pkg/middleware"
)

func PublicRoutes(a *fiber.App) {
	//รายการสินค้า
	services_product := a.Group("/api")
	services_product.Get("/product", middleware.JWTProtected(), service_product.GetProduct)
	services_product.Get("/product/:product_code", middleware.JWTProtected(), service_product.GetProductByID)

	//รายการใบสั้งซื้อ
	services_order := a.Group("/api")
	services_order.Get("/detail_purchaseorder/:order_code", middleware.JWTProtected(), service_order.DetailPurchaseorder)

	//รายละเอียดข้อมูลสมาชิก && รายละเอียดใบสั้งซื้อ ตามสมาชิก
	services_user := a.Group("/api")
	services_user.Get("/detail_member/:member_code", middleware.JWTProtected(), service_member.DetailMebmer)
	services_user.Get("/detail_purchaseorder_bymember/:member_code", middleware.JWTProtected(), service_member.DetailPurchaseorderByMember)
}
