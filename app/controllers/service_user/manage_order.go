package controllers

import (
	queries_member "app/app/queries/service_user"
	"app/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// รายการสินค้า ของสมาชิก
func DetailPurchaseorderByMember(c *fiber.Ctx) error {
	member_code := c.Params("member_code")
	Result := queries_member.GetPurchaseorderByMember(member_code)
	return c.JSON(fiber.Map{
		"result": Result,
		"code":   utils.ResponseCode()["api"]["success"],
		"msg":    utils.ResponseMessage()["api"]["success"],
	})
}
