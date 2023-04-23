package queries

import (
	"app/platform/database"

	struct_order "app/app/models/service_user"
)

// รายการสินค้า ของสมาชิก
func GetPurchaseorderByMember(member_code string) (result []struct_order.DetailOrderByMember) {
	var SqlWhere = ""
	if member_code != "" {
		SqlWhere += " WHERE orders.member_code = '" + member_code + "' "
	}

	sqlStatement := ` SELECT 
							orders.order_code   ,
							members.member_code    ,
							members.name AS member_name ,
							orders.total ,
							(SELECT 
							json_agg(json_build_object ( 'product_code' , MAIN.product_code , 'product_name' , SUB.product_name , 'unit_price' ,SUB.unit_price ,  'qty' ,MAIN.qty , 'total' , SUB.unit_price * MAIN.qty::INT)) AS Product_detail FROM (
								SELECT 
									jsonb_array_elements(product_detail)->>'product_code' AS product_code ,
									jsonb_array_elements(product_detail)->>'qty' AS qty
									FROM Orders 
									WHERE member_code = @member_code 
							) MAIN INNER JOIN products SUB ON MAIN.product_code::INT = SUB.product_code ),
							orders.remark ,
							orders.status ,
							orders.created_at  
						FROM orders 
						INNER JOIN members ON members.member_code = orders.member_code `
	sqlStatement += SqlWhere
	database.DBConn.Raw(sqlStatement, map[string]interface{}{
		"member_code": member_code}).Scan(&result)
	return
}
