
package handler

import (
	"myapp/internal/model"
	"myapp/internal/service"
	"log"
	"github.com/gofiber/fiber/v2"
)

func Transaction(c *fiber.Ctx) error {

	//		get mail from context middleware
	mail := c.Locals("mail")
	if mail == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("UnAuthorized")
	}

	//		get data transfer from http body
	var transaction model.Transaction
	err := c.BodyParser(&transaction)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	//		Transfer Logic	
	TypeTransfer, err := service.TransferLogic(&transaction, mail.(string))
	switch TypeTransfer{
		case "1":
			return c.JSON(model.Response{
				Success: true,
				Message: "Success Transfer",
			})	
		default:
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
}


func BeforeTransfer(c *fiber.Ctx) error {
	//		get mail from context middleware
	mail := c.Locals("mail")
	if mail == nil {	
		return c.Status(fiber.StatusUnauthorized).SendString("Dont have token")
	}

	//		get data transfer from http body
	var transaction model.Transaction
	err := c.BodyParser(&transaction)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("transaction input error")
	}

	//		Logic
	data, typeTransaction, err := service.BeforeTransferLogic(&transaction, mail.(string))	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	switch typeTransaction {
		case "1":
			return c.JSON(model.Response{
				Success: true,
				Message: "Have",
				Data: data,
			})	
		default:
			return c.Status(fiber.StatusInternalServerError).SendString("Dont have action")
	}
}
