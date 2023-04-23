package routes

import (
	"github.com/gofiber/fiber/v2"

	service_order "app/app/controllers/service_order"
	service_product "app/app/controllers/service_product"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	route := a.Group("/api")

	//รายการสินค้า
	route.Get("/product", service_product.GetProduct)
	route.Get("/product/:product_code", service_product.GetProductByID)

	//รายการใบสั้งซื้อ
	route.Get("/detail_purchaseorder/:order_code", service_order.DetailPurchaseorder)

}
