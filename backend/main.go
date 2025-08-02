package main

import (
	"net/http"
	"myapp/internal/handler"
	"myapp/internal/db"
	"log"
)

func main(){

	err := db.Init()
	if err != nil{
		log.Fatalf("Cannt connect DB: %v", err)
	}
	defer db.DB.Close()
	//		Setup route
	http.HandleFunc("/", handler.GetHello)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/register", handler.Register)

	//		Start server
	http.ListenAndServe(":5000", nil)
}