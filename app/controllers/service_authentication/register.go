package controllers

import (
	"app/pkg/utils"

	"github.com/gofiber/fiber/v2"

	struct_authentication "app/app/models/service_authentication"
	queries_authentication "app/app/queries/service_authentication"
)

// สมัครสมาชิก
func Register(c *fiber.Ctx) error {

	//รับค่า & เช็คประเภท
	filter := struct_authentication.Members{}
	if err := c.BodyParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":   utils.ResponseCode()["api"]["invalid_data_type"],
			"msg":    utils.ResponseMessage()["api"]["invalid_data_type"],
			"msglog": err,
		})
	}

	//ไม่อนุญาตค่าว่าง
	if filter.Username == "" || filter.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": utils.ResponseCode()["api"]["parameter_fail"],
			"msg":  utils.ResponseMessage()["api"]["parameter_fail"],
		})
	}

	//เพิ่มข้อมูล
	if err := queries_authentication.Register(filter); !err {
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
