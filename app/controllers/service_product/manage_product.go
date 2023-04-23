package controllers

import (
	"app/pkg/utils"

	struct_product "app/app/models/service_product"
	queries_product "app/app/queries/service_product"

	"github.com/gofiber/fiber/v2"
)

// เพิ่มสินค้า
func CreateProduct(c *fiber.Ctx) error {

	//รับค่า & เช็คประเภท
	filter := struct_product.Products{}
	if err := c.BodyParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":   utils.ResponseCode()["api"]["invalid_data_type"],
			"msg":    utils.ResponseMessage()["api"]["invalid_data_type"],
			"msglog": err,
		})
	}

	//เพิ่มข้อมูล
	if err := queries_product.CreateProduct(filter); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":   utils.ResponseCode()["api"]["cannot_insert"],
			"msg":    utils.ResponseMessage()["api"]["cannot_insert"],
			"msglog": err,
		})
	}

	return c.JSON(fiber.Map{
		"code": utils.ResponseCode()["api"]["success"],
		"msg":  utils.ResponseMessage()["api"]["success"],
	})
}

// รายการสินค้า (ทั้งหมด)
func GetProduct(c *fiber.Ctx) error {
	Result := queries_product.GetProduct("")
	return c.JSON(fiber.Map{
		"result": Result,
		"code":   utils.ResponseCode()["api"]["success"],
		"msg":    utils.ResponseMessage()["api"]["success"],
	})
}

// รายการสินค้า (ตามรหัส)
func GetProductByID(c *fiber.Ctx) error {
	product_code := c.Params("product_code")
	Result := queries_product.GetProduct(product_code)
	return c.JSON(fiber.Map{
		"result": Result,
		"code":   utils.ResponseCode()["api"]["success"],
		"msg":    utils.ResponseMessage()["api"]["success"],
	})
}
