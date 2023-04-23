package routes

import (
	"github.com/gofiber/fiber/v2"

	service_order "app/app/controllers/service_order"
	service_product "app/app/controllers/service_product"
	service_member "app/app/controllers/service_user"
	"app/pkg/middleware"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api")

	//รายการสินค้า
	route.Get("/product", middleware.JWTProtected(), service_product.GetProduct)
	route.Get("/product/:product_code", middleware.JWTProtected(), service_product.GetProductByID)

	//รายการใบสั้งซื้อ
	route.Get("/detail_purchaseorder/:order_code", middleware.JWTProtected(), service_order.DetailPurchaseorder)

	//รายละเอียดข้อมูลสมาชิก
	route.Get("/detail_member/:member_code", middleware.JWTProtected(), service_member.DetailMebmer)

	//รายละเอียดใบสั้งซื้อ ตามสมาชิก
	route.Get("/detail_purchaseorder_bymember/:member_code", middleware.JWTProtected(), service_member.DetailPurchaseorderByMember)
}
