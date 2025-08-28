package handler

import (
	"myapp/internal/model"
	"myapp/internal/service"
	//"myapp/internal/db"	
	"github.com/gofiber/fiber/v2"
)

func GetUserInfo(c *fiber.Ctx) error {	
	//		Get mail from context middleware
	mail := c.Locals("mail")
	if mail == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Dont have token")
	}
	
	//		Get info from DB
	var user model.User
	err := service.GetUserInfoLogic(&user, mail.(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	//		Send data
	return c.JSON(model.Response{
		Success: true,
		Message: "Get Info",
		Data: user,
	})	
}

func GetUserBalance(c *fiber.Ctx) error {
	//		Get mail from context middleware	
	mail := c.Locals("mail")
	if mail == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Dont have token")	
	}

	//		Get u_id from DB
	var user model.User	
	err := service.GetUserBalanceLogic(&user, mail.(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")	
	}	
	//		Send Data
	return c.JSON(model.Response{
		Success: true,
		Message: "Get Balance",
		Data: user,
	})
}