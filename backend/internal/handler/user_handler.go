package handler

import (
	"encoding/json"
	"myapp/internal/model"
	"net/http"
	"myapp/internal/db"	
)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {	
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
	
	//		Get info from DB
	var user model.User
	err := db.DB.QueryRow("select u_name, u_surname, u_mail, u_phone from user where u_mail = ?", mail).Scan(&user.Name, &user.Surname, &user.Mail, &user.Phone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Internal server error",
		})
		return
	}


	//		Send data
	json.NewEncoder(w).Encode(model.Response{
		Success: true,
		Message: "Get Info",
		Data: user,
	})	
	return
}

func GetUserBalance(w http.ResponseWriter, r *http.Request) {
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

	//		Get u_id from DB
	var user model.User	
	err := model.GetUserByEmail(mail.(string), &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Dont have Mail in DB",
		})
		return
	}

	//		Get Balance from DB
	err = model.GetBalance(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Dont have u_id with balance",
		})
		return
	}
	
	//		Send Data
	json.NewEncoder(w).Encode(model.Response{
		Success: true,
		Message: "Get Balance",
		Data: user,
	})
}