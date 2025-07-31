package handler

import (
	"net/http"
	"fmt"
)

func GetHello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello from backend")
}