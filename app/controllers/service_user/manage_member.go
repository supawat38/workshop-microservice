package controllers

import (
	"app/pkg/utils"

	queries_member "app/app/queries/service_user"

	"github.com/gofiber/fiber/v2"
)

// รายการสมาชิก
func DetailMebmer(c *fiber.Ctx) error {
	member_code := c.Params("member_code")
	Result := queries_member.GetMember(member_code)
	return c.JSON(fiber.Map{
		"result": Result,
		"code":   utils.ResponseCode()["api"]["success"],
		"msg":    utils.ResponseMessage()["api"]["success"],
	})
}
