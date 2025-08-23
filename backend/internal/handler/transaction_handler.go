package handler

import (
	"net/http"
	"myapp/internal/model"
	"encoding/json"
	"log"
	"github.com/shopspring/decimal"
)

func Transaction(w http.ResponseWriter, r *http.Request) {
	//		check method
	if r.Method != http.MethodPost{
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Method not allow",
		})
		return
	}

	//		get mail from context middleware
	mail := r.Context().Value("mail");
	if mail == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Dont have token",
		})
		return
	}

	//		get data transfer from http body
	var transaction model.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "transaction input error",
		})
		return
	}

	//		validate input
	err = validate.Struct(transaction)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Validation fail: " + err.Error(),
		})
		return
	}

	amount, err := decimal.NewFromString(transaction.Amount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Invalide amount",
		})
		return
	}

	if amount.Cmp(decimal.NewFromInt(1)) < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "cannt input amount < 1",
		})
		return
	}
	
	//		check type transaction
	if transaction.Type == "1" {
		err = model.Transfer(mail.(string), &transaction)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		//		send data
		json.NewEncoder(w).Encode(model.Response{
			Success: true,
			Message: "Success transfer",
		})	
		return
	}else{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "dont have this action",
		})
		return
	}
}

func BeforeTransfer(w http.ResponseWriter, r *http.Request) {
	//		check method
	if r.Method != http.MethodPost{
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Method not allow",
		})
		return
	}

	//		get mail from context middleware
	mail := r.Context().Value("mail");
	if mail == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Dont have token",
		})
		return
	}

	//		get data transfer from http body
	var transaction model.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "transaction input error",
		})
		return
	}

	if transaction.Type == "1" {
		var data model.CheckUser
		err = model.CheckDataTransfer(mail.(string), transaction.W_id, transaction.Type, transaction.To_w_id, &data)	
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Response{
				Success: false,
				Message: "transaction input error",
			})
			return
		}

		//		send data
		json.NewEncoder(w).Encode(model.Response{
			Success: true,
			Message: "Success transfer",
			Data: data,
		})	
		return
	}else{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "dont have this action",
		})
		return
	}
	
}