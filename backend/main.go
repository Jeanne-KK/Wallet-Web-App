package main

import (
	"myapp/internal/handler"
	"myapp/internal/db"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main(){
	app := fiber.New()
	
	//		Connect DB
	err := db.Init()
	if err != nil{
		log.Fatalf("Cannt connect DB: %v", err)
	}
	defer db.DB.Close()

	//		Setting CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Authorization",
		AllowMethods: "GET, POST",
		AllowCredentials: true,
	}))	

	//		Setup route	
	app.Post("/login", handler.Login)
	app.Post("/register", handler.Register)
	app.Post("/getUserInfo", handler.AuthMiddleware, handler.GetUserInfo)
	app.Post("/getUserBalance", handler.AuthMiddleware, handler.GetUserBalance)
	app.Post("/transfer", handler.AuthMiddleware, handler.Transaction)
	app.Post("/logout", handler.AuthMiddleware, handler.Logout)
	app.Post("/beforetransfer", handler.AuthMiddleware, handler.BeforeTransfer)

	//		Start server
	app.Listen(":5000")
}