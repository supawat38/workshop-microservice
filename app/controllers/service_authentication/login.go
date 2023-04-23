package controllers

import (
	helper "app/app/controllers/common"
	models_authentication "app/app/models/service_authentication"
	queries_authentication "app/app/queries/service_authentication"
	"app/pkg/utils"
	"app/platform/logger"

	"github.com/gofiber/fiber/v2"
)

// เข้าสู่ระบบ
func Login(c *fiber.Ctx) error {

	//รับค่า & เช็คประเภท
	filter := models_authentication.Req_login{}
	if err := c.BodyParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":   utils.ResponseCode()["api"]["invalid_data_type"],
			"msg":    utils.ResponseMessage()["api"]["invalid_data_type"],
			"msglog": err,
		})
	}

	//Validate Struct
	errors := helper.ValidateStruct(filter)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	//Check Member in DB
	Result := queries_authentication.CheckMemberLogin(filter)
	if Result.Username == "" {
		//return HTTP 200 (เพราะเข้าถึงได้ แต่ไม่พบข้อมูล)
		return c.JSON(fiber.Map{
			"code": utils.ResponseCode()["api"]["data_not_found"],
			"msg":  utils.ResponseMessage()["api"]["data_not_found"],
		})
	}

	//ถ้าเจอข้อมูลจะเอาไปสร้าง Token
	var tokenCheck, error = utils.GenerateAccessTokenUser(Result)
	if error != nil {
		logger.SugarLogger.Errorf(error.Error())
		return c.JSON(fiber.Map{
			"code": utils.ResponseCode()["api"]["token_generate_error"],
			"msg":  utils.ResponseMessage()["api"]["token_generate_error"],
		})
	}

	// var Token string
	return c.JSON(fiber.Map{
		"token": tokenCheck,
		"code":  utils.ResponseCode()["api"]["success"],
		"msg":   utils.ResponseMessage()["api"]["success"],
	})
}

// ออกจากระบบ
func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"code": utils.ResponseCode()["api"]["success"],
		"msg":  utils.ResponseMessage()["api"]["success"],
	})
}
