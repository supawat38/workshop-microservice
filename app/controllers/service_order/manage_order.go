package controllers

import (
	"app/pkg/utils"

	struct_order "app/app/models/service_order"
	queries_order "app/app/queries/service_order"

	"github.com/gofiber/fiber/v2"
)

// เพิ่มใบสั่งซื้อ
func CreatePurchaseorder(c *fiber.Ctx) error {

	//รับค่า & เช็คประเภท
	filter := struct_order.ReqOrder{}
	if err := c.BodyParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":   utils.ResponseCode()["api"]["invalid_data_type"],
			"msg":    utils.ResponseMessage()["api"]["invalid_data_type"],
			"msglog": err,
		})
	}

	//เพิ่มข้อมูล
	err, last_order_code := queries_order.CreatePurchaseorder(filter)
	if err != true {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":   utils.ResponseCode()["api"]["cannot_insert"],
			"msg":    utils.ResponseMessage()["api"]["cannot_insert"],
			"msglog": err,
		})
	}

	//อัพเดทราคา total
	status := queries_order.UpdateTotalPurchaseorder(last_order_code)
	if !status {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": utils.ResponseCode()["api"]["cannot_update"],
			"msg":  utils.ResponseMessage()["api"]["cannot_update"],
		})
	}

	return c.JSON(fiber.Map{
		"code": utils.ResponseCode()["api"]["success"],
		"msg":  utils.ResponseMessage()["api"]["success"],
	})
}

// ยกเลิกใบสั่งซื้อ
func CancelPurchaseorder(c *fiber.Ctx) error {

	//รับค่า & เช็คประเภท
	filter := struct_order.ReqOrderCode{}
	if err := c.BodyParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":   utils.ResponseCode()["api"]["invalid_data_type"],
			"msg":    utils.ResponseMessage()["api"]["invalid_data_type"],
			"msglog": err,
		})
	}

	//เช็ตว่ามีในระบบไหม
	Result := queries_order.FindPurchaseorder(filter)
	if Result.OrderCode == 0 {
		//return HTTP 200 (เพราะเข้าถึงได้ แต่ไม่พบข้อมูล)
		return c.JSON(fiber.Map{
			"code": utils.ResponseCode()["api"]["data_not_found"],
			"msg":  utils.ResponseMessage()["api"]["data_not_found"],
		})
	}

	//ยกเลิกใบสั้งซื้อ
	status := queries_order.UpdateStatusPurchaseorder(filter)
	if !status {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": utils.ResponseCode()["api"]["cannot_update"],
			"msg":  utils.ResponseMessage()["api"]["cannot_update"],
		})
	}

	return c.JSON(fiber.Map{
		"code": utils.ResponseCode()["api"]["success"],
		"msg":  utils.ResponseMessage()["api"]["success"],
	})
}

// รายละเอียดใบสั่งซื้อ
func DetailPurchaseorder(c *fiber.Ctx) error {
	order_code := c.Params("order_code")
	Result := queries_order.GetPurchaseorder(order_code)

	return c.JSON(fiber.Map{
		"result": Result,
		"code":   utils.ResponseCode()["api"]["success"],
		"msg":    utils.ResponseMessage()["api"]["success"],
	})
}
