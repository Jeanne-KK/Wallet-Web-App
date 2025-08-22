package main

import (
	"net/http"
	"myapp/internal/handler"
	"myapp/internal/db"
	"log"
	"github.com/rs/cors"
)

func main(){
	mux := http.NewServeMux()
	//		Connect DB
	err := db.Init()
	if err != nil{
		log.Fatalf("Cannt connect DB: %v", err)
	}
	defer db.DB.Close()

	//		Setting CORS
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(mux)


	//		Setup route
	mux.HandleFunc("/", handler.GetHello)
	mux.HandleFunc("/login", handler.Login)
	mux.HandleFunc("/register", handler.Register)
	mux.Handle("/getUserInfo", handler.AuthMiddleware(http.HandlerFunc(handler.GetUserInfo)))
	mux.Handle("/getUserBalance", handler.AuthMiddleware(http.HandlerFunc(handler.GetUserBalance)))
	mux.Handle("/transfer", handler.AuthMiddleware(http.HandlerFunc(handler.Transaction)))
	mux.Handle("/logout", handler.AuthMiddleware(http.HandlerFunc(handler.Logout)))

	//		Start server
	http.ListenAndServe(":5000", cors)
}