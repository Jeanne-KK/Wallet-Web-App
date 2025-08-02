package model

type InputLogin struct{
	Mail string `json:"mail" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type InputRegister struct {
	Mail string `json:"mail" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Name string `json:"name" validate:"required,alpha,max=100"`
	Surname string `json:"surname" validate:"required,alpha,max=100"`
	Phone string `json:"phone" validate:"required,numeric,min=9,max=10"`
}