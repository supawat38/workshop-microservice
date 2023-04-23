package queries

import (
	struct_order "app/app/models/service_order"
	"app/platform/database"
	"time"
)

// เพิ่มใบสั่งซื้อ
func CreatePurchaseorder(Parameter struct_order.ReqOrder) (success bool, result_order_code uint) {

	// วันที่ปัจจุบัน
	dTimeNow := time.Now().Local()
	dTimeString := dTimeNow.Format("2006-01-02 15:04:05")

	//เพิ่มข้อมูลตารางออเดอร์
	sqlStatement := ` INSERT INTO Orders
						( member_code , status , product_detail , remark , created_at) `
	sqlStatement += ` 	VALUES (
							@member_code ,
							'success' ,
							@product_detail ,
							@remark ,
							@created_at
						) `
	if database.DBConn.Exec(sqlStatement,
		map[string]interface{}{
			"member_code":    Parameter.MemberCode,
			"product_detail": Parameter.ProductDetail,
			"remark":         Parameter.Remark,
			"created_at":     dTimeString,
		}).Error != nil {
		return
	}

	//เอารหัส order ล่าสุด
	sqlLastCode := ` SELECT order_code FROM Orders ORDER BY order_code DESC LIMIT 1 `
	database.DBConn.Raw(sqlLastCode).Scan(&result_order_code)

	success = true
	return
}

// อัพเดทราคา total
func UpdateTotalPurchaseorder(OrderCode uint) (status bool) {
	queryUpdate := `UPDATE Orders SET total = (	
						SELECT SUM(SUB.unit_price * MAIN.qty::INT) AS total FROM (
							SELECT 
								jsonb_array_elements(product_detail)->>'product_code' AS product_code ,
								jsonb_array_elements(product_detail)->>'qty' AS qty
								FROM Orders 
								WHERE order_code = @order_code 
						) MAIN INNER JOIN products SUB ON MAIN.product_code::INT = SUB.product_code) WHERE order_code = @order_code`
	err := database.DBConn.Exec(queryUpdate, map[string]interface{}{
		"order_code": OrderCode}).Error
	if err != nil {
		status = false
		return
	}
	status = true
	return
}

// ตรวจสอบก่อน
func FindPurchaseorder(Parameter struct_order.ReqOrderCode) (Result struct_order.ReqOrderCode) {

	var tQueryWhere = ""
	tQueryWhere += ` order_code = @order_code AND member_code =  @member_code `

	//อัพเดทข้อมูลที่ตารางผู้ใช้
	query_result := `SELECT order_code FROM Orders WHERE `
	query_result += tQueryWhere
	database.DBConn.Raw(query_result, map[string]interface{}{
		"order_code":  Parameter.OrderCode,
		"member_code": Parameter.MemberCode}).Scan(&Result)
	return Result
}

// ยกเลิกใบสั่งซื้อ
func UpdateStatusPurchaseorder(Parameter struct_order.ReqOrderCode) (status bool) {
	queryUpdate := `UPDATE Orders SET status = 'cancel' WHERE order_code = @order_code AND member_code = @member_code`
	err := database.DBConn.Exec(queryUpdate, map[string]interface{}{
		"order_code":  Parameter.OrderCode,
		"member_code": Parameter.MemberCode}).Error
	if err != nil {
		status = false
		return
	}
	status = true
	return
}

// รายการใบสั้งซื้อ (ตามไอดี)
func GetPurchaseorder(order_code string) (result struct_order.DetailOrder) {
	var SqlWhere = ""
	if order_code != "" {
		SqlWhere += " WHERE order_code = @order_code "
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
									WHERE order_code = @order_code 
							) MAIN INNER JOIN products SUB ON MAIN.product_code::INT = SUB.product_code ),
							orders.remark ,
							orders.status ,
							orders.created_at  
						FROM orders 
						INNER JOIN members ON members.member_code = orders.member_code `
	sqlStatement += SqlWhere
	database.DBConn.Raw(sqlStatement, map[string]interface{}{
		"order_code": order_code}).Scan(&result)
	return
}
