package queries

import (
	"app/platform/database"

	struct_product "app/app/models/service_product"
)

// เพิ่มข้อมูลสินค้า
func CreateProduct(product struct_product.Products) error {
	return database.DBConn.Create(&product).Error
}

// ข้อมูลสินค้า
func GetProduct(product_code string) (result []struct_product.Products) {
	var SqlWhere = ""
	if product_code != "" {
		SqlWhere += " WHERE product_code = '" + product_code + "' "
	}

	sqlStatement := ` SELECT * FROM Products `
	sqlStatement += SqlWhere
	database.DBConn.Raw(sqlStatement).Scan(&result)
	return
}
