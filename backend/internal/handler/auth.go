package handler

import (
	"encoding/json"
	"net/http"
	"myapp/internal/model"
	"log"
	"github.com/go-playground/validator/v10"
	"myapp/internal/utils"
	"myapp/internal/db"	
	"context"
	"myapp/internal/service"
)

var validate = validator.New()

func Login(w http.ResponseWriter, r *http.Request){
	//		Check method
	if r.Method != http.MethodPost{
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Method not allow",
		})
		return
	}

	//		Check format json
	var data model.InputLogin	
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Invalid JSON format",
		})
		return
	}

	//		Validate input
	err = validate.Struct(data)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Validation fail: " + err.Error(),
		})
		return
	}
	//log.Printf("recieve mail: %s, password %s", data.Mail, data.Password)

	//		Check user from mail
	var user model.InputLogin
	err = db.DB.QueryRow("select u_mail, u_password from user where u_mail = ?", data.Mail).Scan(&user.Mail, &user.Password)
	if err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Invalid email or password",	
		})	
		return
	}
	//log.Println("have user from mail")

	//		Check password
	if !utils.CheckPassword(data.Password, user.Password){
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Invalid email or password",	
		})	
		return
	}

	//		Create JWT token
	token, err := utils.GenerateToken(user.Mail)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Fail to create token",	
		})	
		return
	}
	
	//		Set cookie
	utils.SetCookie(w, token)

	//		return token
	json.NewEncoder(w).Encode(model.Response{
		Success: true,
		Message: "Login success",	
	})
	
}

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

func AuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println("FROM middleware")

		//		Get cookie
		cookie, err := utils.GetCookie(r, "jwt")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.Response{
				Success: false,
				Message: "Cookie not found",
			})
		return
		}
	
		data, err := utils.ValidateToken(cookie)
		//log.Println(data.Mail)
		ctx := context.WithValue(r.Context(), "mail", data.Mail)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
