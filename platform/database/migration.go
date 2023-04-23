package database

import (
	service_authentication "app/app/models/service_authentication"
	service_order "app/app/models/service_order"
	service_product "app/app/models/service_product"
)

func Init() {
	DBConn.AutoMigrate(
		&service_authentication.Members{},
		&service_product.Products{},
		&service_order.Orders{},
	)
}
