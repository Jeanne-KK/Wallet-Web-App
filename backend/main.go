package main

import (
	"net/http"
	"myapp/internal/handler"
)

func main(){

	//		Setup route
	http.HandleFunc("/", handler.GetHello)

	//		Start server
	http.ListenAndServe(":5000", nil)
}