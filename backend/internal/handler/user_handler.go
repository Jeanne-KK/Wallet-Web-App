package handler

import (
	"myapp/internal/model"
	"myapp/internal/db"	
	"github.com/gofiber/fiber/v2"
)

func GetUserInfo(c *fiber.Ctx) error {	

	//		Get mail from context middleware
	mail := c.Context().Value("mail")
	if mail == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Dont have token")
	}
	
	//		Get info from DB
	var user model.User
	err := db.DB.QueryRow("select u_name, u_surname, u_mail, u_phone from user where u_mail = ?", mail).Scan(&user.Name, &user.Surname, &user.Mail, &user.Phone)
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
	mail := c.Context().Value("mail")
	if mail == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Dont have token")
		
	}

	//		Get u_id from DB
	var user model.User	
	err := model.GetUserByEmail(mail.(string), &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Dont have Mail in DB")
	}

	//		Get Balance from DB
	err = model.GetBalance(&user)
	if err != nil {	
		return c.Status(fiber.StatusInternalServerError).SendString("Dont have u_id with balance")
	}
	
	//		Send Data
	return c.JSON(model.Response{
		Success: true,
		Message: "Get Balance",
		Data: user,
	})
}