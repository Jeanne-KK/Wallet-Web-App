package handler

import (
	"myapp/internal/model"
	"log"
	"github.com/go-playground/validator/v10"
	"myapp/internal/utils"
	"myapp/internal/service"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func Login(c *fiber.Ctx) error {
	//		Check method
	if c.Method() != fiber.MethodPost{
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method not allowed")
	}

	//		Check format json
	var data model.InputLogin	
	err := c.BodyParser(&data)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}		

	//		Logic login
	token, err := service.LoginUser(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	
	//		Set cookie
	utils.SetCookie(c, token)

	//		return token
	return c.JSON(model.Response{
		Success: true,
		Message: "Login success",
	})	
}

/*
func Register(w http.ResponseWriter, r *http.Request){
	
	//		Check method
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//		Check format json
	var data model.InputRegister	
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil{
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//		logic Register
	token, err := service.RegisterUser(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//		Use httpCookie
	utils.SetCookie(w, token)

	//		return token
	json.NewEncoder(w).Encode(model.Response{
		Success: true,
		Message: "Register success",
	})
}

func Logout(w http.ResponseWriter, r *http.Request){
	//		Check method
	if r.Method != http.MethodPost{
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Method not allow",
		})
		return
	}

	//		Get mail from context middleware
	mail := r.Context().Value("mail")
	if mail == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Dont have token",
		})
		return
	}

	//		Clear cookie
	utils.ClearCookie(w)

	//		Response
	json.NewEncoder(w).Encode(model.Response{
		Success: true,
		Message: "Logout success",
	})
}
*/

func AuthMiddleware(c *fiber.Ctx) error{
	log.Println("FROM middleware")

	//		Get cookie
	cookie, err := utils.GetCookie(c, "jwt")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Cookie not found"})
	}	
	data, err := utils.ValidateToken(cookie)
	//log.Println(data.Mail)
	c.Locals("mail", data.Mail)
	
	return c.Next()
}

