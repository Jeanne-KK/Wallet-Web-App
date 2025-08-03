package handler

import (
	"encoding/json"
	"net/http"
	"myapp/internal/model"
	"log"
	"github.com/go-playground/validator/v10"
	"myapp/internal/utils"
	"myapp/internal/db"	
	"database/sql"
	"context"
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
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Method not allow",
		})
	}

	//		Check format json
	var data model.InputRegister	
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Invalid JSON format",
		})
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

	//		Check have mail already	
	var user model.InputLogin
	err = db.DB.QueryRow("select u_mail from user where u_mail = ?", data.Mail).Scan(&user)
	//log.Println(user)
	log.Println(user.Mail)
	if err != nil {
		if err != sql.ErrNoRows {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Email already exists",	
		})	
		return
		}
	}

	//		Hash password
	hash, err := utils.HashPassword(data.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Fail to hash password",	
		})	
		return
	}
	//log.Println(hash)

	//		Make new pass with hash
	newUser := model.InputRegister{
		Mail: data.Mail,
		Password: hash,
		Name: data.Name,
		Surname: data.Surname,
		Phone: data.Phone,
	}

	_, err = db.DB.Exec("insert into user (u_mail, u_password, u_name, u_surname, u_phone) values (?, ?, ?, ?, ?)", newUser.Mail, newUser.Password, newUser.Name, newUser.Surname, newUser.Phone)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Fail to create user",	
		})	
		return
	}

	//		Create JWT token
	token, err := utils.GenerateToken(newUser.Mail)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Fail to create token",	
		})	
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
