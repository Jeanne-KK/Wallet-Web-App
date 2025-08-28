package model

import (
	"myapp/internal/db"
	"database/sql"
)

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

type User struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Mail string `json:"mail"`
	Phone string `json:"phone"`
	UserId string `json:"u_id"`
	WalletId string `json:"w_id"`
	Balance string `json:"w_balance"`	
}

func GetUserByEmail(mail string,user *User) (error) {
	err := db.DB.QueryRow("select u_id, u_mail from user where u_mail = ?", mail).Scan(&user.UserId, &user.Mail)
	if err != nil {
		return err
	}
	return nil
}

func GetBalance(user *User) (error){
	err := db.DB.QueryRow("select w_id, w_balance from wallet w, user u where w.u_id = u.u_id and u.u_id = ?", user.UserId).Scan(&user.WalletId, &user.Balance)
	if err != nil {
		return err
	}
	return nil
}

func CheckMailExist(mail string) (bool, error){
	var hold string
	err := db.DB.QueryRow("select u_mail from user where u_mail = ?", mail).Scan(&hold)
	if err == sql.ErrNoRows{
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func CreateUser(newUser InputRegister) (error){
	_, err := db.DB.Exec("insert into user (u_mail, u_password, u_name, u_surname, u_phone) values (?, ?, ?, ?, ?)", newUser.Mail, newUser.Password, newUser.Name, newUser.Surname, newUser.Phone)
	return err
}

func GetPassHash(data string) (string, error) {
	var passHash string
	err := db.DB.QueryRow("select u_password from user where u_mail = ?", data).Scan(&passHash)
	if err != nil {
		return "", err
	}
	return passHash, nil
}

func GetInfo(user *User, mail string) (error){
	err := db.DB.QueryRow("select u_name, u_surname, u_mail, u_phone from user where u_mail = ?", mail).Scan(&user.Name, &user.Surname, &user.Mail, &user.Phone)
	if err != nil {
		return err
	}
	return nil
}